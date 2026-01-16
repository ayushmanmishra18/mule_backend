package audit

import (
	"context"
	"time"
	"muleshield/pkg/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Logger struct {
	svc   *dynamodb.Client
	table string
}

func NewLogger(region, table string) *Logger {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	return &Logger{svc: dynamodb.NewFromConfig(cfg), table: table}
}

func (l *Logger) LogHighRisk(tx models.Transaction, risk models.RiskResponse) {
	item, _ := attributevalue.MarshalMap(map[string]interface{}{
		"log_id":         tx.TransactionID,
		"timestamp":      time.Now().Format(time.RFC3339),
		"risk_score":     risk.RiskScore,
		"recommendation": risk.Recommendation,
	})
	l.svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(l.table),
		Item:      item,
	})
}