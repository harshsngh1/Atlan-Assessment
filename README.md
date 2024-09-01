# Atlan-Assessment

High Level Diagram : https://drive.google.com/file/d/1sOf1Ci82FtcyL5ejVYlZss40f4EtI6zf/view?usp=sharing
## Description
This project consists of a set of microservices designed to integrate with Kafka for processing and notifying metadata changes. The core components include:
- **Monte Carlo Ingestion Service:** This service receives metadata from Monte Carlo via HTTP POST requests, processes the data, and publishes it to a Kafka topic.
- **Compliance Service:** This service reads from a Kafka topic, processes metadata related to PII (Personally Identifiable Information) and GDPR (General Data Protection Regulation), and notifies external systems about data access control and compliance.

## Modules
1. Monte Carlo Ingestion Service
    - Purpose: To ingest metadata from Monte Carlo and push it to Kafka.
    - Key Components:
        - main.go :  Entry point of the service.
        - ingest.go :  Function to ingest metadata from Monte Carlo.
        - kafka.go :  Function to publish metadata to Kafka.
2. Compliance Service
    - Purpose: To read PII and GDPR annotations from Kafka and notify external systems.
    - Key Components:
        - main.go :  Entry point of the service.
        - compliance.go :  Function to process PII and GDPR annotations.

## Technology Used
- Go: Programming language used for the services.
- Kafka: Messaging system for handling metadata and notifications.
- Zookeeper: Coordination service for Kafka.
- Docker: Containerization platform to manage services and dependencies.
- Docker Compose: Tool to define and run multi-container Docker applications.

## Steps to Setup
1. Clone the Repository
```
git clone <repository-url>
cd <repository-directory>
```
2. Build the Docker Images  
Navigate to the directories of each service and build the Docker images:
```
cd monte-carlo-ingestion
docker build -t monte-carlo-ingestion .

cd ../compliance-service
docker build -t compliance-service .
```
3. Start Services Using Docker Compose  
Ensure you are in the directory containing docker-compose.yml and run:
```
docker-compose up
```
This will start Kafka, Zookeeper, Monte Carlo Ingestion Service, and Compliance Service.
4. Create Kafka Topics  
Create the necessary Kafka topics manually if they are not created automatically:
```
docker exec -it <kafka-container-id> kafka-topics --create --topic monte-carlo-metadata --partitions 1 --replication-factor 1 --zookeeper zookeeper:2181
docker exec -it <kafka-container-id> kafka-topics --create --topic pii-gdpr-annotations --partitions 1 --replication-factor 1 --zookeeper zookeeper:2181
```
5. To Test This Setup
    1. Test Monte Carlo Ingestion Service  
    Use curl to send a test POST request:
    ```
    curl -X POST http://localhost:8080/ingest -H "Content-Type: application/json" -d '{"key": "value"}'
    ```
    This should result in the data being published to the monte-carlo-metadata Kafka topic.
    2. Test Compliance Service  
    Ensure the compliance service is running and has successfully started. The compliance service should automatically process messages from the pii-gdpr-annotations topic and log notifications about data access control.  
    Check the logs for any errors or confirmation that the notifications are being processed:
    ```
    docker logs <compliance-service-container-id>
    ```

## Troubleshooting
- Kafka Issues: Check the Kafka and Zookeeper logs if topics are not being created or if there are leader election issues.
```
docker logs <kafka-container-id>
docker logs <zookeeper-container-id>
```
- Service Logs: Inspect logs of Monte Carlo Ingestion and Compliance Service for errors or connection issues.
```
docker logs <monte-carlo-ingestion-container-id>
docker logs <compliance-service-container-id>
```