# 测试用例模块

## 测试用例数据格式
    {
        "name": "用例名称",
        "type": "用例类型",
        "priority": "用例优先级",
        "description": "用例描述",
        "url": "用例文件（支持文件格式由对应框架决定）"
    }

## 测试用例生成流程

~~~ 
    测试用例信息填写 --> 上传用例文件 --> 使用对应测试框架调试 --> 调试通过 --> 确认添加后生成yaml环境文件保存
~~~ 


# 任务格式
    
    {
        "name": "任务名称",
        "type": "任务类型：实时任务/定时任务",
        "description": "任务描述（optional）",
        "rank": "任务优先级（紧急程度: 抢占式执行？）",
        "status": "任务状态（生命周期管理）",
        "testcaseSuite": "关联用例集"
    }

# 任务流程

~~~ 
    Task --> testcaseSuite --> testcase --> url --> file --> run
    
    Worker --> Environment Doctor through Configuration file --> Accept testcase
    
    Conf
~~~ 

# 配置文件（Configuration）

~~~ 
    文件类型：yaml
    
    tool: environment doctor
~~~ 