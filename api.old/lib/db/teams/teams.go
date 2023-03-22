package team

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)



type TeamService struct {
	client *dynamodb.Client
	table  string
}

func NewTeamService(table string) (*TeamService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	_, err = client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: &table,
	})

	if err != nil {
		// If table doesn't exist, create it
		if strings.Contains(err.Error(), "ResourceNotFoundException") {
			_, err = client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
				AttributeDefinitions: []types.AttributeDefinition{
					{
						AttributeName: aws.String("id"),
						AttributeType: types.ScalarAttributeTypeS,
					},
				},
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("id"),
						KeyType:       types.KeyTypeHash,
					},
				},
				TableName: &table,
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &TeamService{
		client: client,
		table:  table,
	}, nil
}

func (s *TeamService) Create(team *Team) error {
	_, err := s.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: team.ID},
			"name": &types.AttributeValueMemberS{Value: team.Name},
			"members": &types.AttributeValueMemberL{
				Value: make([]types.AttributeValue, len(team.Members)),
			},
		},
		TableName: &s.table,
	})
	if err != nil {
		return err
	}

	for i, member := range team.Members {
		_, err = s.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName: &s.table,
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: team.ID},
			},
			ExpressionAttributeNames: map[string]string{
				"#members": "members",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":val": &types.AttributeValueMemberS{Value: member.ID},
			},
			UpdateExpression: aws.String("SET #members[" + fmt.Sprint(i) + "] = :val"),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TeamService) Get(id string) (*Team, error) {
	res, err := s.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: &s.table,
	})
	if err != nil {
		return nil, err
	}

	if res.Item == nil {
		return nil, errors.New("team not found")
	}

	members := make([]Member, 0)
	for _, member := range res.Item["member"].(*types.AttributeValueMemberL).Value {
		members = append(members, Member{
			ID: member.(*types.AttributeValueMemberS).Value,
		})
	}

	return &Team{
		ID:      res.Item["id"].(*types.AttributeValueMemberS).Value,
		Name:    res.Item["name"].(*types.AttributeValueMemberS).Value,
		Members: members,
	}, nil
}
