package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/innocentcivilian/interviewtask/helpers"
	"github.com/innocentcivilian/interviewtask/service"
	"github.com/innocentcivilian/interviewtask/util"
)

var usecase service.DeviceService

// type book struct {
// 	ID    int    `dynamodbav:"id"`
// 	Title string `dynamodbav:"title"`
// }
type handler struct {
	usecase service.DeviceService
}

// Get a single user
func (h *handler) Get(ctx context.Context, id string) (helpers.Response, error) {
	device, err := h.usecase.Get(ctx, id)
	if err != nil {
		return util.ResponseMessage(util.InternalError, http.StatusInternalServerError)
	}
	return util.ResponseData(util.OK, device, http.StatusOK)
}

// Create a user
func (h *handler) Create(ctx context.Context, body []byte) (helpers.Response, error) {
	user := &users.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	if err := h.usecase.Create(ctx, user); err != nil {
		return helpers.Fail(err, http.StatusInternalServerError)
	}

	return helpers.Success(user, http.StatusCreated)
}
func main() {
	usecase, err := service.Init()
	if err != nil {
		log.Panic(err)
	}
	// lambda.Start(handler)
	h := &handler{usecase}
	lambda.Start(helpers.Router(h))

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
	// lambda.Start(handler)
}
