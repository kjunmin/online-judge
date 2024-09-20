package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	pb "github.com/kjunmin/online-judge/server/common/api"
)

type store struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func NewStore(tableName string) *store {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("localhost"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://dynamodb-local:8000"}, nil
			},
		)),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "abcd", SecretAccessKey: "a1b2c3", SessionToken: "",
				Source: "Mock credentials for local instance",
			},
		}),
	)
	if err != nil {
		panic(err)
	}
	c := dynamodb.NewFromConfig(cfg)
	return &store{
		DynamoDbClient: c,
		TableName:      tableName,
	}
}

func (s *store) TableExists(ctx context.Context) (bool, error) {
	exists := true
	_, err := s.DynamoDbClient.DescribeTable(
		ctx, &dynamodb.DescribeTableInput{TableName: aws.String(s.TableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", s.TableName)
		} else {
			log.Printf("Could not determine the existence of table %v, Error: %v", s.TableName, err)
		}
		exists = false
	}
	return exists, err
}

func (s *store) CreateProblemsTable(ctx context.Context) (*types.TableDescription, error) {
	var tableDesc *types.TableDescription
	table, err := s.DynamoDbClient.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("ProblemID"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("Title"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("ProblemID"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("Title"),
			KeyType:       types.KeyTypeRange,
		}},
		TableName: aws.String(s.TableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(3),
			WriteCapacityUnits: aws.Int64(3),
		},
	})
	if err != nil {
		log.Printf("Couldn't create table %v. Error %v\n", s.TableName, err)
	} else {
		waiter := dynamodb.NewTableExistsWaiter(s.DynamoDbClient)
		err = waiter.Wait(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(s.TableName)}, 5*time.Minute)
		if err != nil {
			log.Printf("Wait for table exists failed. Error: %v", err)
		}
		tableDesc = table.TableDescription
	}
	return tableDesc, err
}

func (s *store) AddProblem(ctx context.Context, problem *pb.Problem) error {
	item, err := attributevalue.MarshalMap(problem)
	if err != nil {
		log.Printf("Marshal Error %v\n", err)
		panic(err)
	}
	_, err = s.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.TableName), Item: item,
	})
	if err != nil {
		log.Printf("Unable to add item to table. Error %v\n", err)
	}
	return err
}
