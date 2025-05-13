

官方教程：https://go-zero.dev/docs/tutorials/cli/docker


1、使用gin写几个接口

文件详细看main.go


2、使用 goctl docker 命令创建 Dockerfile 文件

    ```goctl docker --go main.go --exe api_docker```

其中，main.go 是包含 main 方法的入口文件

api_docker 是要创建的镜像名称


3、使用 Dockerfile 文件创建镜像文件

如果还没有安装docker，请看 [CentOS8.2版安装docker【20250509】](https://www.yuque.com/nhjclxc/java/qdo45feyriwdrmly?singleDoc#)

在 api_docker 项目目录下，执行 docker build 命令，生成镜像
```docker build -t api_docker:v1.0 .```

其中，
- -t是 --tag的缩写，
- api_docker 是镜像名称
- :v1.0 表示镜像版本
- api_docker 后面一个参数是镜像生成路径，"." 表示在当前目录下生成

4、启动容器

这条命令的作用是：运行一个容器，基于你构建好的 my-gin-app 镜像，以后台方式启动，并将容器的端口映射到主机端口。

```docker run -d -p 8080:8080 --name api_docker-app api_docker```


| 部分                      | 含义                                                                             |
|-------------------------| ------------------------------------------------------------------------------ |
| `docker run`            | 启动一个新的容器                                                                       |
| `-d`                    | **Detached mode**，即“后台运行容器”。如果不用这个，容器会在前台运行，终端会被占用                             |
| `-p 8099:8080`          | **端口映射**：将宿主机（本机）的 `8099` 端口映射到容器的 `8080` 端口（格式是 `宿主机:容器`）<br>这通常是 gin 应用监听的端口 |
| `--name api_docker-app` | 给这个容器起个名字叫 `api_docker-app`，便于后续用命令管理（比如 `docker stop api_docker-app`）                       |
| `api_docker`            | 要运行的镜像名。这里指的是你之前用 `docker build -t api_docker .` 创建的镜像                         |


其他常用命令：
| 命令                           | 作用                          |
| ---------------------------- | --------------------------- |
| `docker ps`                  | 查看当前运行的容器                   |
| `docker stop gin-app`        | 停止容器                        |
| `docker start gin-app`       | 启动已停止的容器                    |
| `docker logs gin-app`        | 查看容器日志（你可以排查运行中 gin 的错误信息）  |
| `docker exec -it gin-app sh` | 进入容器内部（需要容器使用的是带 shell 的镜像） |










