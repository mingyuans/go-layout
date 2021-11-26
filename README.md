## How to run docker in local machine

1- Build image: 
```shell
make image
```

2- Run docker 
```shell
docker run -p {your_local_port}:8080 -v {your_local_config_dir_path}:/etc/iam  {image}
```
