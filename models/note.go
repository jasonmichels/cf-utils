package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Note struct {
	ID        string `json:"id"`
	ClaimID   string `json:"claimId"`
	Body      string `json:"body"`
	CreatorID string `json:"creatorId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func (n *Note) ConvertToDynamodbAttributes() (map[string]types.AttributeValue, error) {
	// Create a map using JSON field names
	noteMap := map[string]interface{}{
		"id":        n.ID,
		"claimId":   n.ClaimID,
		"body":      n.Body,
		"creatorId": n.CreatorID,
		"createdAt": n.CreatedAt,
		"updatedAt": n.UpdatedAt,
	}

	av, err := attributevalue.MarshalMap(noteMap)
	if err != nil {
		return nil, err
	}

	return av, nil
}
