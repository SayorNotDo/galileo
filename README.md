# 基于go-kratos框架开发的自动化测试分析服务

## 项目结构

```.
├── api
│   ├── core
│   ├── engine
│   ├── file
│   ├── management
│   └── user
├── app
│   ├── core
│   ├── engine
│   ├── file
│   ├── management
│   └── user
├── database
├── deploy
├── docs
├── ent
│   ├── enttest
│   ├── group
│   ├── hook
│   ├── migrate
│   ├── predicate
│   ├── project
│   ├── runtime
│   ├── schema
│   ├── task
│   ├── testcase
│   ├── testcasesuite
│   └── user
└── pkg
    ├── ctxdata
    ├── encryption
    ├── errResponse
    ├── responseEncoder
    ├── transport
    └── utils
```

传统的软件测试工作在于执行与交付，保证软件的质量，我设想的是能够通过大数据的实现对于测试的分析，从而驱动软件的质量提高

基于大数据解决方案思考如下：

1.数据收集：收集大规模的测试数据，包括测试行为数据、系统日志、传感器数据等，确保数据完整、真实，以便评估产品或系统的性能和效果。

2.数据清洗和处理：对收集到的原始数据进行清洗和处理，以去除错误、重复或不完整的数据。包括：数据去重、异常值处理、数据格式标准化等，以确保数据的准确性和一致性

3.数据分析和探索：使用数据分析技术和工具，对清洗后的数据进行探索和分析。包括数据可视化、统计分析、机器学习等，以发现其中的模式、趋势和洞察，为后续的测试评估提供依据

4.测试评估和优化：基于数据分析的结果，设计测试方案和评估指标，以评估产品或系统的性能、可靠性和用户体验等方面。根据评估结果，进行优化和改进，以提升产品或系统的效果和竞争力。

5.故障排除和优化：通过分析大数据，识别和分析可能的故障和性能瓶颈，采取相应的措施；来排除故障和改进性能

6.预测分析和决策支持，基于大数据分析的结果，使用预测模型和算法，对未来的趋势和情景进行预测，帮助做出更有效的决策，并优化产品或系统的发展策略

以上6个步骤是初步构思的解决方案，下面详细叙述实现的想法

### 数据采集

数据采集作为解决方案的首要步骤，涉及到收集大规模的实际数据，以供后续的数据清洗、处理和分析使用。其需要进行以下几方面的考虑：

定义数据目标，数据采集之前需要明确数据需求和目标，确定想要收集的数据类型、字段、时间范围等。

选择数据源，确定要从哪些数据源中采集数据，如系统日志、数据库、传感器、第三方API等

数据获取方式，根据数据源的类型和数据访问方式，选择适当的方式来获取数据，目前初步设想是通过自研的SDK接入测试框架，收集测试时的各类数据，通过HTTP/RPC等方式进行上报

确定采集范围，定义采集的数据范围和时间窗口

采集频率和时间计划

数据存储和管理

数据安全和隐私

数据质量验证

### 数据清洗

### 数据分析

### 流程草图

### 流程简述：

开发测试数据采集SDK，示例如通过Python开发，实现数据采集与上报功能，接入现有的自动化测试框架等，在测试执行时调用上报

上报数据通过HTTP/RPC等，可进行鉴权校验以拦截非法数据，接口对上报数据进行清洗与处理后写入kafka消费队列

大数据使用Doris列式存储系统，通过（Routine Load）例行导入功能，实时从指定的Kafka处读取数据，导入Doris中

后台通过实现SQL生成器，向Doris发起SQL查询，返回结果

前端基于结果生成可视化报表

#### Python SDK开发

##### 具体功能：

##### 数据采集

##### 数据上报

##### 采集指标

    网速、FPS？、测试耗时（具体看测试用例）、响应时间（针对于接口？）

#### 后台数据入库

测试数据通过HTTP/RPC接口上报

接口需要经过网关接口鉴权(确保数据来源)，同时收集收集请求的相关信息，转化为服务端属性

ETL 代表提取、转换和加载，是组织将多个系统中的数据组合到单个数据库、数据存储空间、数据仓库或数据湖中的传统方法

### 数仓理论

#### 概念

#### 特点

#### 架构分层

#### 元数据

## 测试管理模块

### 项目
### 任务
### 用例/用例集