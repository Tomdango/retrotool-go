package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseParameters struct {
	Endpoint string
	Port     string
	Name     string
	Username string
	Password string
}

func GetDatabaseParameters(cfg aws.Config) DatabaseParameters {
	ssmClient := ssm.NewFromConfig(cfg)

	params, err := ssmClient.GetParametersByPath(context.TODO(), &ssm.GetParametersByPathInput{
		Path:           aws.String("/prod/db/"),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("Unable to retrieve SSM DB parameters, %v", err)
	}

	paramMap := make(map[string]string, len(params.Parameters))
	for _, parameter := range params.Parameters {
		name := strings.TrimPrefix(*parameter.Name, "/prod/db/")
		paramMap[name] = *parameter.Value
	}

	return DatabaseParameters{
		Endpoint: paramMap["endpoint"],
		Username: paramMap["username"],
		Password: paramMap["password"],
		Name:     paramMap["name"],
		Port:     paramMap["port"],
	}
}

func NewConnection(cfg aws.Config) (*gorm.DB, error) {
	var dsn string
	if os.Getenv("ENVIRONMENT") == "local" {
		dsn = "host=localhost user=local password=local dbname=postgres sslmode=disable"
	} else {
		params := GetDatabaseParameters(cfg)
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=verify-full",
			params.Endpoint,
			params.Username,
			params.Password,
			params.Name,
			params.Port,
		)
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
