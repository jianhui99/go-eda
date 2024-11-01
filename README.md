### Event-Driven Architecture (EDA) with RabbitMQ

```mermaid
graph LR
    subgraph EDA
        A[User Service] --> Exchange
        Exchange --> B[Email Service]
        B[Email Service] --> Exchange
        Exchange --> C[Log Service]
    end
```

### RabbitMQ
set up RabbitMQ with docker
```
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:2.11-management
```

create new user and grant permission, and remove default guest user
```
docker exec rabbitmq rabbitmqctl add_user [user_name] [password]
docker exec rabbitmq rabbitmqctl set_user_tags [user_name] administrator
docker exec rabbitmq rabbitmqctl delete_user guest
```

create vhost and grant permissions
```
docker exec rabbitmq rabbitmqctl add_vhost [name]
docker exec rabbitmq rabbitmqctl set_permissions -p [vhost_name] [user_name] ".*" ".*" ".*"
```


declare exchanges and set permissions
```
docker exec rabbitmq rabbitmqadmin  declare exchange --vhost=[vhost_name] name=[exchange_name]  type=topic -u [user_name] -p [password] durable=true 
docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers jian customer_events "^customers.*" "^customers.*"
```