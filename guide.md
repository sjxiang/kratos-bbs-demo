



### bbs demo


环境准备
    go、protoc 以及相关插件


安装 kratos 命令行工具
    go install github.com/go-kratos/kratos/cmd/kratos/v2@latest（科学上网）


创建项目模板
    kratos new kratos-bbs-demo

    cd kratos-bbs-demo/
    go mod tidy
    make all
    kratos run