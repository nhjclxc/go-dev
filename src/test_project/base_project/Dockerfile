FROM registry.xxx.net:9191/infra/busybox:stable-glibc-amd64
WORKDIR /app

RUN uname -m
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY ./bin/base_project /app/base_project
RUN chmod 777 /app/base_project
