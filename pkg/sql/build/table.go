package build

import (
	"matrixone/pkg/sql/op"
	"matrixone/pkg/sql/op/relation"
	"matrixone/pkg/sql/tree"
)

func (b *build) buildTable(stmt *tree.TableName) (op.OP, error) {
	if len(stmt.SchemaName) == 0 {
		return b.getTable(b.db, string(stmt.ObjectName))
	}
	return b.getTable(string(stmt.SchemaName), string(stmt.ObjectName))
}

func (b *build) getTable(schema string, name string) (op.OP, error) {
	db, err := b.e.Database(schema)
	if err != nil {
		return nil, err
	}
	r, err := db.Relation(name)
	if err != nil {
		return nil, err
	}
	return relation.New(name, schema, r), nil
}
