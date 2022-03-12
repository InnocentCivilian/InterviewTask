## Dynamodb setup
in order to setup aws dynamodb setup locally create a `docker-compose.yml` with following contents:
```yml
version: '3.7'
services:
  dynamodb:
    image:  dockerhub.ir/amazon/dynamodb-local
    container_name: my-dynamodb
    hostname: dynamodb
    networks:
      - backend
    restart: always
    volumes:
      -  ./my-dynamodb-data:/home/dynamodblocal/data
    ports:
      - 8000:8000
    command: "-jar DynamoDBLocal.jar -sharedDb -dbPath /home/dynamodblocal/data/"
networks:
  backend:
    name: hmn-sam-backend
```
and run `docker-compose up -d` to start it.
## Run sam
run `sam local start-api --docker-network hmn-sam-backend` in project root directory
## Build project
run `sam build` in project root directory
