FROM registry.xxx.net:9191/infra/busybox:stable-glibc-amd64
WORKDIR /app

RUN uname -m
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY ./bin/go_base_project /app/go_base_project
RUN chmod 777 /app/go_base_project
