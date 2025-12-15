


`vim kafka_server_jaas.conf`
`vim docker-compose.yaml`

`docker-compose down && docker-compose up -d`

创建kafka接入用户

```shell

# 创建用户
  
docker exec kafka /opt/kafka/bin/kafka-configs.sh \
  --bootstrap-server kafka:9092 \
  --alter \
  --add-config 'SCRAM-SHA-256=[password=pwd123]' \
  --entity-type users \
  --entity-name user1
  
#重启kafka
docker-compose up -d kafka

#查看用户

docker exec kafka /opt/kafka/bin/kafka-configs.sh \
  --bootstrap-server kafka:9092 \
  --entity-type users \
  --entity-name user1 \
  --describe
  
```

`docker logs -f kafka`