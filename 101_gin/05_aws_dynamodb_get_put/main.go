package main

// AWS SDK for Go: https://docs.aws.amazon.com/sdk-for-go/
// AWS SDK for Go (DynamoDB): https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html
// dynamo is an expressive DynamoDB client for Go, https://github.com/guregu/dynamo
// go get -u github.com/guregu/dynamo

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"log"
)

type ShortUrl struct {
	// filed name must start with a capital letter
	Id  string `json:"id" dynamo:"id"`
	Url string `json:"url" dynamo:"url"`
	Ttl int64  `json:"ttl" dynamo:"ttl"`
}

// Error{409, "resource_conflict", dao}
type ApiError struct {
	Code  int         `json:"code"`
	Type  string      `json:"type"`
	More  interface{} `json:"more"`
	Error error       `json:"-"`
}

func main() {
	engine := gin.Default()
	engine.POST("/ddb", handleRequest)
	engine.Run()

	/*
		curl --location --request POST 'http://localhost:8080/ddb' \
		--header 'Content-Type: application/json;charset=utf8' \
		--data-raw '{
		    "id": "0E24A",
		    "url": "https://star.ettoday.net/news/1827711?redirect=1",
		    "ttl": 1602236766
		}'
	*/
}

func handleRequest(ctx *gin.Context) {
	var shortUrl ShortUrl
	if ctx.ShouldBind(&shortUrl) == nil {
		log.Println(shortUrl)
	}

	var dao ShortUrl
	dao, err := createShortUrl(shortUrl)
	if err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(200, dao)
	}
}

func createShortUrl(shortUrl ShortUrl) (ShortUrl, *ApiError) {
	// better to use ~/.aws/credentials, more detail: https://docs.aws.amazon.com/zh_tw/cli/latest/userguide/cli-configure-profiles.html
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		fmt.Println("Fail to create AWS session", err)
	}
	db := dynamo.New(sess)
	table := db.Table("tmc-stg-short_url")

	// get item
	var dao ShortUrl
	err = table.Get("id", shortUrl.Id).One(&dao)
	if err != nil {
		log.Println("Fail to get item by id from DynamoDB table", shortUrl.Id, err)
	}

	// check conflict
	if &dao != nil {
		return dao, &ApiError{
			Code:  409,
			Type:  "resource_conflict",
			More:  dao,
			Error: errors.New("unavailable"),
		}
	}
	// put item
	// TODO: ConditionCheck to handle race condition: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TransactWriteItems.html
	err = table.Put(shortUrl).Run()
	if err != nil {
		log.Println("Fail to put item to DynamoDB table", err)
	}

	return shortUrl, nil
}
