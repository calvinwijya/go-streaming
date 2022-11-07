# streaming pipeline

This service is for Real-Time Streaming Pipeline Engine for Storing Hundreds Million of Data
It follows Clean Architecture.
It has below layers:

Delivery Layer: It deals with consuming events from Kafka <br>
Schema:

```
type Delivery interface {
	ConsumeEvents(topicName string)
}
```

<br>

Usecase Layer: It deals with processing the Kafka Events and invoking Datastore Repository <br>
Schema:

```
type Usecase interface {
	ProcessData(events []interface{})
}
```

<br>

Repository Layer: It deals with storing the processed data to Cassandra <br>

<br>

Schema:

```
type Repository interface {
	Store(data []map[string]interface{}) (err error)
}
```

Sample Events from Kafka would be of below format <br>

```
{
    "pid": 123,
    "recommended_pids": [456,789]
}
```

Note: These exercises are based upon the presentation slides

Command to connect to docker container of Kafka: docker exec -it -w /opt/bitnami/kafka kafka-server /bin/sh
Command to stop clients: Ctrl+C

Step 1 (Kafka):

1. Create Kafka Topic
   - bin/kafka-topics.sh --create --topic streaming-pipeline --bootstrap-server localhost:9092
   - to check: bin/kafka-topics.sh --describe --topic streaming-pipeline --bootstrap-server localhost:9092
2. Check all Kafka Topics
   - bin/kafka-topics.sh --bootstrap-server=localhost:9092 --list
3. Publish Sample Message
   - bin/kafka-console-producer.sh --broker-list localhost:9092 --topic streaming-pipeline
   - type "hello world"
4. Consume/check message
   - bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic streaming-pipeline

===================================================================================================================

Command to connect to docker container of Cassandra: docker exec -it cassandra /bin/sh

Step 2 (Cassandra):

1. Connect to DB
   - cqlsh
2. Setup Keyspace
   - CREATE KEYSPACE streaming_pipeline WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'} AND durable_writes = true;
3. Create Table schema
   - CREATE TABLE streaming_pipeline.pipeline_data (
     pid bigint,
     recommended_pids list<bigint>,
     PRIMARY KEY (pid)
     );
4. Basic DML operations
   - Insert: INSERT INTO streaming_pipeline.pipeline_data (pid , recommended_pids ) VALUES ( 1,[2,3,4,5,6,7,8,9,10]);
   - Select: SELECT \* FROM streaming_pipeline.pipeline_data WHERE pid=1;

===================================================================================================================

Command to connect to docker container of Kafka: docker exec -it -w /opt/bitnami/kafka kafka-server /bin/sh
Command to connect to docker container of Cassandra: docker exec -it cassandra /bin/sh
Command to stop clients: Ctrl+C

Step 3: (Pipeline Flow)

1. Open kafka producer and consumer
   - bin/kafka-console-producer.sh --broker-list localhost:9092 --topic streaming-pipeline
   - bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic streaming-pipeline
2. Run the pipeline
   - go build
   - ./go-streaming
3. Push message to kafka
   - { "pid": 123, "recommended_pids": [456,789] }
   - { "pid": 101, "recommended_pids": [102,103] }
4. Check the data in cassandra
   - Select: SELECT \* FROM streaming_pipeline.pipeline_data WHERE pid=123;
