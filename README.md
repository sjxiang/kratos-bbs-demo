

## bbs，基于 kratos 实现，教程


```text
01 项目创建
    搭个脚手架

02 API 定义与生成
    xiaohongshu.proto 小红书论坛，集美打卡 (: 
    复制粘贴 UnimplementedXiaohongshuServer 下面的方法

03 数据库接入与配置修改
    讲真，配置这块儿，比 go-zero 强，不是黑盒（要自己翻代码才知道写的啥）


04 项目结构与依赖注入
    
    通过 google wire 实现，解耦 

        service/
            NewXiaohongshuService(?)
                biz.UserUseCase
        
        biz/
            NewUserUseCase(?)
                UserRepo (interface)  # 拼装实体以及通过编译器确保方法都实现
                ProfileRepo

        data/
            NewUserRepo(?) && NewProfileRepo(?)
                (Data)

            NewData(?)
                DB

            NewDB(?)
                gorm
        
        config/
            database.dsn
        

    但也很难构思这么巧妙，依赖也很难搞



05 biz 层开发

    填充业务逻辑


06 自定义中间件

    JWT 鉴权 

    这套框架最狗屎的问题，gRPC 插件适配 (:


07 cors 与 HTTP 中间件自定义
    跨域 浏览器


08 HTTP 错误返回结构
    略

09 data 层开发
    

10 错误处理
    略

11 构建和部署
    找个现成的前端联调下

```
