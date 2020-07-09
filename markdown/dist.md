# dist源码分析

dist命令在一个独立目录中，没有依赖，编译通过GOROOT_BOOTSTRAP导入进来的低版本go命令进行编译

## dist命令作用
* `./cmd/dist/dist env -p` 输出env环境变量
* `./cmd/dist/dist bootstrap -a` 安装编译命令