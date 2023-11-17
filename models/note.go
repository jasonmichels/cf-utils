package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Note struct {
	ID        string `json:"id"`
	ClaimID   string `json:"claimId"`
	Body      string `json:"body"`
	CreatorID string `json:"creatorId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func (n *Note) ConvertToDynamodbAttributes() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"id":        {S: aws.String(n.ID)},
		"claimId":   {S: aws.String(n.ClaimID)},
		"body":      {S: aws.String(n.Body)},
		"creatorId": {S: aws.String(n.CreatorID)},
		"createdAt": {N: aws.String(fmt.Sprintf("%d", n.CreatedAt))},
		"updatedAt": {N: aws.String(fmt.Sprintf("%d", n.UpdatedAt))},
	}
}
