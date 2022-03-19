## How to run docker in local machine

1- Build image: 
```shell
make image
```

2- Run docker 
```shell
docker run -p {your_local_port}:8080 -v {your_local_config_dir_path}:/etc/iam  {image}
```

## 性能监控

1- pprof

服务默认开启 pprof, 默认访问路径为: http://{ip}/debug/pprof
