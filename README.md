

## bbs，基于 kratos 实现

### 

项目创建
    搭个脚手架

API 定义与生成
    xiaohongshu.proto 小红书论坛，集美打卡 (: 
    复制粘贴 UnimplementedXiaohongshuServer 下面的方法


数据库接入与配置修改
    讲真，配置这块儿，比 go-zero 强，不是黑盒，要自己翻代码才知道写的啥


项目结构与依赖注入（通过 google wire 实现）

    解耦 

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



biz 层开发和中间件

自定义中间件

cors 与 HTTP 中间件自定义

HTTP 错误返回结构

data 层开发

错误处理

构建和部署


###