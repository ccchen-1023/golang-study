package main

// Go implementation of JSON Web Tokens: https://github.com/dgrijalva/jwt-go
// JWT debugger: https://jwt.io/

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/lestrrat-go/jwx/jwk"
	"log"
	"net/http"
	"strings"
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
	engine.POST("/cognito", handleRequest)
	engine.Run()

	/*
		curl --location --request POST 'http://localhost:8080/cognito' \
		--header 'Content-Type: application/json;charset=utf8' \
		--header 'Authorization: Bearer eyJraWQiOiJMZTZnVU5TTmlUaTl0bU5zY...' \
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

	token, er := auth(ctx.Request)
	if token == nil || token.Valid == false || er != nil {
		ctx.JSON(403, ApiError{
			Code: 403,
			Type: "unauthorized",
		})
	} else {
		var dao ShortUrl
		dao, err := createShortUrl(shortUrl)
		if err != nil {
			ctx.JSON(err.Code, err)
		} else {
			ctx.JSON(200, dao)
		}
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func auth(req *http.Request) (*jwt.Token, error) {

	tokenString := extractToken(req)

	// AWS Cognito public keys are available at address:
	// https://cognito-idp.{region}.amazonaws.com/{userPoolId}/.well-known/jwks.json
	publicKeysURL := "https://cognito-idp.us-west-2.amazonaws.com/us-west-2_FzFbL2FDR/.well-known/jwks.json"

	// Start with downloading public keys information
	// The .Fetch method is used from https://github.com/lestrrat-go/jwx package
	keySet, err := jwk.Fetch(publicKeysURL)
	if err != nil {
		log.Printf("failed to parse key: %s", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodRS256.Name {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		keys := keySet.LookupKeyID(kid)
		if len(keys) == 0 {
			return nil, fmt.Errorf("key %v not found", kid)
		}

		var raw interface{}
		return raw, keys[0].Raw(&raw)
	})
	log.Println(token)
	// TODO: new auth group to replace with classifier
	groups := token.Claims.(jwt.MapClaims)["cognito:groups"]
	authGroup := false
	for _, group := range groups.([]interface{}) {
		if group == "classifier" {
			authGroup = true
		}
	}
	if authGroup == false {
		token.Valid = false
	}
	return token, nil
}

func extractAccessToken(req *http.Request) string {
	bearToken := req.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
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
