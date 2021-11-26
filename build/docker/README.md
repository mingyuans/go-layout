## 注意事项

docker 目录下的子目录需要跟 cmd/ 保持一致；

比如 cmd/iam-apisever, 那么，docker 这里也需要是 docker/iam-apiserver/Dockerfil.
这是因为脚本会自动根据 Dockerfile 所在的文件夹名触发对应的 cmd 产品编译。如果 Dockerfile 这里的文件夹名与 cmd 不对应，会导致产物编译失败。
