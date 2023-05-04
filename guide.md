



### bbs demo


环境准备
    go、protoc 以及相关插件


安装 kratos 命令行工具
    go install github.com/go-kratos/kratos/cmd/kratos/v2@latest（科学上网）


创建项目
    # 使用默认模板创建项目
    kratos new kratos-bbs-demo
    
    # 进入项目目录
    cd kratos-bbs-demo

    # 拉取项目依赖
    go mod download

    # 插件(protoc 相关)
    make init


代码生成和运行

    生成
        # 生成所有proto源码、wire等等
        go generate ./...

    运行
        # 运行项目
        kratos run

        # 输出
        INFO msg=config loaded: config.yaml format: yaml # 默认载入 configs/config.yaml 配置文件
        INFO msg=[gRPC] server listening on: [::]:9000 # gRPC服务监听 9000 端口
        INFO msg=[HTTP] server listening on: [::]:8000 # HTTP服务监听 8000 端口

        