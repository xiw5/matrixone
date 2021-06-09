package build

import (
	"fmt"
	"matrixone/pkg/container/types"
	"matrixone/pkg/sql/colexec/aggregation"
	"matrixone/pkg/sql/colexec/aggregation/avg"
	"matrixone/pkg/sql/colexec/aggregation/count"
	"matrixone/pkg/sql/colexec/aggregation/max"
	"matrixone/pkg/sql/colexec/aggregation/min"
	"matrixone/pkg/sql/colexec/aggregation/starcount"
	"matrixone/pkg/sql/colexec/aggregation/sum"
	"matrixone/pkg/sql/op"
	"matrixone/pkg/sql/op/projection"
	"matrixone/pkg/sql/op/summarize"
	"matrixone/pkg/sql/tree"
)

var AggFuncs map[string]int = map[string]int{
	"avg":   aggregation.Avg,
	"max":   aggregation.Max,
	"min":   aggregation.Min,
	"sum":   aggregation.Sum,
	"count": aggregation.Count,
}

func (b *build) hasSummarize(ns tree.SelectExprs) bool {
	for _, n := range ns {
		if e, ok := n.Expr.(*tree.FuncExpr); ok {
			if name, ok := e.Func.FunctionReference.(*tree.UnresolvedName); ok {
				if _, ok = AggFuncs[name.Parts[0]]; ok {
					return true
				}
			}

		}
	}
	return false
}

func (b *build) buildSummarize(o op.OP, ns tree.SelectExprs) (op.OP, error) {
	var es []aggregation.Extend

	{
		for _, n := range ns {
			if _, ok := n.Expr.(*tree.FuncExpr); !ok {
				return nil, fmt.Errorf("noaggregated column '%s'", n.Expr)
			}
		}
	}
	o, err := b.stripSummarize(o, ns)
	if err != nil {
		return nil, err
	}
	for _, n := range ns {
		expr := n.Expr.(*tree.FuncExpr)
		name, ok := expr.Func.FunctionReference.(*tree.UnresolvedName)
		if !ok {
			return nil, fmt.Errorf("illegal expression '%s'", n)
		}
		op, ok := AggFuncs[name.Parts[0]]
		if !ok {
			return nil, fmt.Errorf("unimplemented aggregated functions '%s'", name.Parts[0])
		}
		alias := string(n.As)
		switch e := expr.Exprs[0].(type) {
		case *tree.NumVal:
			if len(alias) == 0 {
				alias = "count(*)"
			}
			agg, err := newAggregate(op, types.Type{Oid: types.T_int64, Size: 8})
			if err != nil {
				return nil, err
			}
			es = append(es, aggregation.Extend{
				Agg:   agg,
				Alias: alias,
				Op:    aggregation.StarCount,
			})
		case *tree.UnresolvedName:
			if len(alias) == 0 {
				alias = fmt.Sprintf("%s(%s)", name.Parts[0], e.Parts[0])
			}
			typ, ok := o.Attribute()[e.Parts[0]]
			if !ok {
				return nil, fmt.Errorf("unknown column '%s' in aggregation", e.Parts[0])
			}
			agg, err := newAggregate(op, typ)
			if err != nil {
				return nil, err
			}
			es = append(es, aggregation.Extend{
				Op:    op,
				Agg:   agg,
				Alias: alias,
				Name:  e.Parts[0],
			})

		}
	}
	return summarize.New(o, es), nil
}

func (b *build) stripSummarize(o op.OP, ns tree.SelectExprs) (op.OP, error) {
	var es []*projection.Extend

	for _, n := range ns {
		if _, ok := n.Expr.(*tree.FuncExpr).Exprs[0].(*tree.NumVal); !ok {
			e, err := b.buildExtend(o, n.Expr.(*tree.FuncExpr).Exprs[0])
			if err != nil {
				return nil, err
			}
			es = append(es, &projection.Extend{E: e})
			n.Expr.(*tree.FuncExpr).Exprs[0] = &tree.UnresolvedName{
				Parts: [4]string{e.String()},
			}
		}
	}
	if len(es) > 0 {
		return projection.New(o, es), nil
	}
	return o, nil
}

func newAggregate(op int, typ types.Type) (aggregation.Aggregation, error) {
	switch op {
	case aggregation.Avg:
		switch typ.Oid {
		case types.T_float32, types.T_float64:
			return avg.NewFloat(typ), nil
		case types.T_int8, types.T_int16, types.T_int32, types.T_int64:
			return avg.NewInt(typ), nil
		case types.T_uint8, types.T_uint16, types.T_uint32, types.T_uint64:
			return avg.NewUint(typ), nil
		}
	case aggregation.Max:
		switch typ.Oid {
		case types.T_int8:
			return max.NewInt8(typ), nil
		case types.T_int16:
			return max.NewInt16(typ), nil
		case types.T_int32:
			return max.NewInt32(typ), nil
		case types.T_int64:
			return max.NewInt64(typ), nil
		case types.T_uint8:
			return max.NewUint8(typ), nil
		case types.T_uint16:
			return max.NewUint16(typ), nil
		case types.T_uint32:
			return max.NewUint32(typ), nil
		case types.T_uint64:
			return max.NewUint64(typ), nil
		case types.T_float32:
			return max.NewFloat32(typ), nil
		case types.T_float64:
			return max.NewFloat64(typ), nil
		case types.T_char, types.T_varchar:
			return max.NewStr(typ), nil
		}
	case aggregation.Min:
		switch typ.Oid {
		case types.T_int8:
			return min.NewInt8(typ), nil
		case types.T_int16:
			return min.NewInt16(typ), nil
		case types.T_int32:
			return min.NewInt32(typ), nil
		case types.T_int64:
			return min.NewInt64(typ), nil
		case types.T_uint8:
			return min.NewUint8(typ), nil
		case types.T_uint16:
			return min.NewUint16(typ), nil
		case types.T_uint32:
			return min.NewUint32(typ), nil
		case types.T_uint64:
			return min.NewUint64(typ), nil
		case types.T_float32:
			return min.NewFloat32(typ), nil
		case types.T_float64:
			return min.NewFloat64(typ), nil
		case types.T_char, types.T_varchar:
			return min.NewStr(typ), nil
		}
	case aggregation.Sum:
		switch typ.Oid {
		case types.T_float32, types.T_float64:
			return sum.NewFloat(typ), nil
		case types.T_int8, types.T_int16, types.T_int32, types.T_int64:
			return sum.NewInt(typ), nil
		case types.T_uint8, types.T_uint16, types.T_uint32, types.T_uint64:
			return sum.NewUint(typ), nil
		}
	case aggregation.Count:
		return count.New(typ), nil
	case aggregation.StarCount:
		return starcount.New(typ), nil
	}
	return nil, fmt.Errorf("unimplemented aggregation '%v' for '%s'", op, typ)
}
