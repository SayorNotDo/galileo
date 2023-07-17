# 测试用例模块

## 测试用例数据格式（json）
    {
        "framework": "测试框架选择",
        "steps": "用例步骤",
        "expected_result": "预期结果",
        "data": "用例数据",
    }

## 在线测试用例编写

```
    用例编写 --> 生成指定格式用例文件 --> 上传文件存储
```

# 测试框架加载
~~~
    设定框架路径（远程or本地） --> docker部署环境 --> 加载配置文件 --> 回调等待执行
~~~ 

## 用例Debug流程

    {
       "name": "用例名称",
        "type": "用例类型", 
        "content": "用例内容"
    }

## 测试用例上传

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