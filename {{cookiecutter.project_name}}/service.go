package __cookiecutter_project_name__

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	log "github.com/sirupsen/logrus"
	"os"
)
type ObjectService interface {
	SaveObject(ctx context.Context, createObject *CreateObject) (Object, error)}

func Init(integration bool) (ObjectService, error) {

	region := os.Getenv("AWS_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		log.Info("error connecting to dynamodb")
		return nil, err
	}
	ddb := dynamodb.New(sess)
	if integration == false {
		xray.Configure(xray.Config{LogLevel: "trace"})
		xray.AWS(ddb.Client)
	}
	tableName := os.Getenv("TABLE_NAME")
	repository := NewDynamoDBRepository(ddb, tableName)
	usecase :=  &Usecase{Repository: repository}
	return usecase, nil
}
