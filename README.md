# Apache Kafka Example on Go

## Run Apache Kafka

To run:

```sh
$ MY_IP=your-ip docker-compose up
```

If we want to create a topic, like `acl-chat` you must run:

```sh
$ docker run --net=host --rm confluentinc/cp-kafka:5.3.0 kafka-topics --create --topic acl-chat --partitions 3 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```
$ docker run --net=host --rm confluentinc/cp-kafka:5.3.0 kafka-topics --create --topic acl-chat --partitions 3 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
This command creates a topic named `acl-chat` with `4 partitions` and `replication factor of 2`.

## Run the server

The default host is `http://localhost:8080`, you can change this configuration on `.env` file

```sh
make server-run
```

## Run the clients

You can run as many clients as you want, executing this:

```sh
make client-run
```