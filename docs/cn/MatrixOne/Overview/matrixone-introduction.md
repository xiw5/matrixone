# **MatrixOne简介** 

MatrixOne是一款面向未来的超融合异构云原生数据库，通过超融合数据引擎支持事务/分析/流处理等混合工作负载，通过异构云原生架构支持跨机房协同/多地协同/云边协同。简化开发运维，消简数据碎片，打破数据的系统、位置和创新边界。

## **核心特性** 
### **超融合引擎**
* **超融合引擎**
     
     融合数据引擎，单数据库即可支持TP、AP、时序、机器学习等混合工作负载。

* **内置流引擎** 
     
     利用独有的增量物化视图能力，无需跨数据库即可实现实时数据流处理。


### **异构云原生**

* **异构统一**
     
     支持跨机房协同/多地协同/云边协同，实现无感知扩缩容，提供高效统一的数据管理。

* **多地多活**
  
     MatrixOne采用最优的一致性协议，实现业内最短网络延迟的多地多活。


### **极致的性能**

* **高性能**

     特有的因子化计算和向量化执行引擎，支持极速的复杂查询。单表、星型和雪花查询都具备极速分析性能。


* **强一致**
  
     提供跨存储引擎的高性能全局分布式事务能力，在保证极速分析性能的同时支持更新、删除和实时点查询。

* **高可用**
  
     存算分离，支持存储节点与计算节点独立扩缩容，高效应对负载变化。


## **用户价值**

* **简化数据开发和运维**
  
     随着业务发展，企业使用的数据引擎和中间件越来越多，而每一个数据引擎平均依赖5+个基础组件，存储3+个数据副本，每一个数据引擎都要各自安装、监控、补丁和升级。这些都导致数据引擎的选型、开发及运维成本高昂且不可控。在MatrixOne的一体化架构下，用户使用单个数据库即可服务多种数据应用，引入的数据组件和技术栈减少80%，大大简化了数据库管理和维护的成本。


* **消减数据碎片和不一致**
  
     在既有复杂的系统架构内，存在多条数据管道多份数据存储冗余。数据依赖复杂，导致数据更新维护复杂，上下游数据不一致问题频发，人工校对难度增大。MatrixOne的高内聚架构和独有的增量物化视图能力，使得下游可以支持上游数据的实时更新，摆脱冗余的ETL流程，实现端到端实时数据处理。


* **无需绑定基础设施**
  
     因为基础设施的碎片化，企业的私有化数据集群和公有云数据集群之间数据架构和建设方案割裂，数据迁移成本高。而数据上云一旦选型确定数据库厂商，后续的集群扩容、其他组件采购等都将被既有厂商绑定。MatrixOne提供统一的云边基础架构和高效统一的数据管理，企业数据架构不再被基础设施绑定，实现单数据集群跨云无感知扩缩容，提升性价比。


* **极速的分析性能**
  
     目前，由于缓慢的复杂查询性能以及冗余的中间表，数据仓库在业务敏捷性上的表现不尽人意，大量宽表的创建也严重影响迭代速度。MatrixOne通过特有的因子化计算和向量化执行引擎，支持极速的复杂查询，单表、星型和雪花查询都具备极速分析性能。

* **像TP一样可靠的AP体验**
  
     传统数据仓库数据更新代价非常高，很难做到数据更新即可见。在营销风控，无人驾驶，智能工厂等实时计算要求高的场景或者上游数据变化快的场景中，当前的大数据分析系统无法支持增量更新，往往需要做全量的更新，耗时耗力。MatrixOne通过提供跨存储引擎的高性能全局分布式事务能力，支持条级别的实时增量更新，在保证极速分析性能的同时支持更新、删除和实时点查询。

* **不停服自动扩缩容**
  
     传统数仓无法兼顾性能和灵活度，性价比无法做到最优。MatrixOne基于存算分离的技术架构，支持存储节点与计算节点独立扩缩容，高效应对负载变化。


## **相关信息**

本节描述了MatrixOne的产品核心特性和用户价值主张，如果您想了解有关MatrixOne产品架构设计的更多信息，请参阅以下内容：  
* [MatrixOne技术架构](matrixone-architecture.md)