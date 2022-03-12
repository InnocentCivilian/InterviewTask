package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	dto "github.com/innocentcivilian/interviewtask/dto/createdevicerequest"
	"github.com/innocentcivilian/interviewtask/util"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var data dto.CreateDeviceRequest
	//parse request body into dto
	err := json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       util.InvalidJsonError,
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	errMsg, errValidation := util.Validate(data)

	if errValidation != nil {
		//validation failed
		var resp = dto.ValidationError{
			Message: util.InputDataInvalid,
			Errors:  errMsg,
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       "internal error",
				StatusCode: http.StatusInternalServerError,
			}, nil
		}
		return events.APIGatewayProxyResponse{
			Body:       string(data),
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	// sess := session.Must(session.NewSession())
	// db := dynamo.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})
	// table := dynamodb.Put()
	// table.Put(device).Run()
	fmt.Println(device.Id)
	av, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Devices"),
	}
	fmt.Printf("%s", av)
	_, err = db.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s %s", err, err.Error())
	}

	// dynamodb.Put(dynamodb.Put{TableName: "Devices",Item: &device})
	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: http.StatusOK,
	}, nil
	// response.StatusCode = http.StatusOK
	// response.Body = string(data)

	// return response, nil
}

// type book struct {
// 	ID    int    `dynamodbav:"id"`
// 	Title string `dynamodbav:"title"`
// }

func main() {

	// validate = validator.New()
	// sess := session.Must(session.NewSession())
	// db := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})

	// deviceTable := db.CreateTable("Books", dynamo.AWSEncoding(input)).Run()
	// if deviceTable != nil {
	// 	// log.Fatalf("Got error calling CreateTable: %s", deviceTable)
	// }
	// params := &dynamodb.CreateTableInput{
	// 	TableName: aws.String("ProductCatalog"),
	// 	AttributeDefinitions: []*dynamodb.AttributeDefinition{
	// 		{
	// 			AttributeName: aws.String("Id"),
	// 			AttributeType: aws.String("N"),
	// 		},
	// 	},
	// 	KeySchema: []*dynamodb.KeySchemaElement{
	// 		{
	// 			AttributeName: aws.String("Id"),
	// 			KeyType:       aws.String("HASH"),
	// 		},
	// 	},
	// 	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	// 		ReadCapacityUnits:  aws.Int64(10),
	// 		WriteCapacityUnits: aws.Int64(5),
	// 	},
	// }
	// r, err := db.CreateTable(params)
	// fmt.Println(r)
	// fmt.Println(err)
	// err := db.Table("Books").Put(dynamo.AWSEncoding(book{
	// 	ID:    42,
	// 	Title: "Principia Discordia",
	// })).Run()
	// fmt.Println(err)
	// // When getting an item you MUST pass a pointer to AWSEncoding!
	// var someBook book
	// fmt.Println(err)

	// err2 := db.Table("Books").Get("ID", 555).One(dynamo.AWSEncoding(&someBook))
	// fmt.Println(err2)
	// fmt.Println(someBook.Title)
	// tableName := "Devices"
	// sess := session.Must(session.NewSession())
	// svc := dynamo.New(sess, &aws.Config{Endpoint: aws.String("http://dynamodb:8000")})

	// deviceTable := svc.CreateTable(tableName, input)
	// // if deviceTable != nil {
	// // 	log.Fatalf("Got error calling CreateTable: %s", deviceTable)
	// // }

	// fmt.Println("Created the table", deviceTable)
	// fmt.Println("Created the table", deviceTable.Run())
	// fmt.Println("Created the table", deviceTable)
	lambda.Start(handler)
}
