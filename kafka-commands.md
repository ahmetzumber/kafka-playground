# Kafka Commands

## STEP 1: Get Kafka
```zsh
tar -xzf kafka_2.13-3.2.1.tgz
cd kafka_2.13-3.2.1
```

## STEP 2: Start The Kafka Environment
We started to run seperate two service for Kafka as called Zookeeper and Kafka broker service.

```zsh
# Start the ZooKeeper service
$ bin/zookeeper-server-start.sh config/zookeeper.properties
```
Open another terminal session and run:
```zsh
# Start the Kafka broker service
$ bin/kafka-server-start.sh config/server.properties
```

## STEP 3: Create a Topic
Open another terminal session and run to create a topic.
```zsh
$ bin/kafka-topics.sh --create --topic ${topic_name} --bootstrap-server localhost:9092
```

## STEP 4: Write Some Events into the Topic
```zsh
$ bin/kafka-console-producer.sh --topic ${topic_name} --bootstrap-server localhost:9092
This is my first event
This is my second event
```
You can stop the event producing with `CTRL+C`.

## STEP 5: Reading The Events
Open another terminal session and run the console consumer client to read the events you just created.
```zsh
$ bin/kafka-console-consumer.sh --topic ${topic_name} --from-beginning --bootstrap-server localhost:9092
```





### References 
https://kafka.apache.org/quickstart























