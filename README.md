## Pre requirments
* [Docker](https://www.docker.com/products/docker-desktop)
* [Aws CLI](https://aws.amazon.com/cli)
* [SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-windows.html)
* [Go lang](https://go.dev/dl/)
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
Note that `dockerhub.ir` is a mirror
and run `docker-compose up -d` to start it.(don't forget to have sam installed and run `aws configure` to setup aws credentials)
Then Apply the [datastore](infrastructure/datastore.yml) with CloudFormation or call `Migrate()` function within [migration](infrastructure/migration.go)
## Run sam
run `sam local start-api --docker-network hmn-sam-backend` in project root directory where `hmn-sam-backend` is the network name defined in `docker-compose` configuration.
## Build project
run `sam build` in project root directory
## Run Tests
run `go test ./..` after a successful build to run all tests (or `go build && go test ./..`). 
* both build and testing stages happen every time a push or merge happens on `master` branch

## Code Architecture
this code tries to follow clean archiecture and groups logic and implemention related to an intedend domain into it's own zone
busniness logic and border between trusted and untrusted context is implemented in use case
## Files
* `.github` holds github action for continuous integration and continuous delivery with build and test stages,depolyment also can be added to this action file
* `cmd` holds executables that can be build and put into aws cloud
* `infrastructure` holds database schema in 2 forms `datastore.yml` and `dynamodbconfig.go` which resolves dynamodb config based on running environment
* `dto` holds data transfare objects(both request and response) and validation rules
* `helpers` holds routing logic and request dispacher
* `model` holds plain old go object (!) that describes the model
* `service` holds both internal and external (exportable) services provided by the application
* `util` holds general use case functions and consts to be used in other layers (message codes , response template and so on)
* `template.yml` holds general use case functions and consts to be used in other layers (message codes , response template and so on)
#Deployment
* Create a stack in AWS cloudformation
* Deploy to AWS using SAM `sam deploy --guided`


#Doc Refrences
* [dynamodb golang SDK](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
* [aws dynamodb cli](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)
