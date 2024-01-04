package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Client struct {
	ClientID    string `json:"clientId"`
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (c *Client) ConvertToDynamodbAttributes() (map[string]types.AttributeValue, error) {
	tempClientMap := map[string]interface{}{
		"clientId":    c.ClientID,
		"name":        c.Name,
		"contact":     c.Contact,
		"phone":       c.Phone,
		"description": c.Description,
		"createdAt":   c.CreatedAt,
		"updatedAt":   c.UpdatedAt,
	}

	av, err := attributevalue.MarshalMap(tempClientMap)
	if err != nil {
		return nil, err
	}

	return av, nil
}
