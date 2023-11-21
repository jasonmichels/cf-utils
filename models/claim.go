package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Claim struct {
	ClaimID               string  `json:"claimId"`
	ClaimNumber           string  `json:"claimNumber"`
	Insured               string  `json:"insured"`
	ClaimCost             float32 `json:"claimCost"`
	ClosingCost           float32 `json:"closingCost"`
	Status                string  `json:"status"`
	Description           string  `json:"description"`
	ClientID              string  `json:"clientId"`
	AssignedUserID        string  `json:"assignedUserId"`
	AssignedUserFirstName string  `json:"assignedUserFirstName"`
	AssignedUserLastName  string  `json:"assignedUserLastName"`
	CreatorID             string  `json:"creatorId"`
	CreatorFirstName      string  `json:"creatorFirstName"`
	CreatorLastName       string  `json:"creatorLastName"`
	WorkflowStep          string  `json:"workflowStep"`
	ContactPerson         string  `json:"contactPerson"`
	ClaimantAddress       string  `json:"claimantAddress"`
	DateOfLoss            int64   `json:"dateOfLoss"`
	ClaimantName          string  `json:"claimantName"`
	ClaimantContact       string  `json:"claimantContact"`
	ClaimantPhone         string  `json:"claimantPhone"`
	ClaimantEmail         string  `json:"claimantEmail"`
	ContactPhone          string  `json:"contactPhone"`
	ContactEmail          string  `json:"contactEmail"`
	ClientBillingMethod   string  `json:"clientBillingMethod"`
	BillablePercentage    float32 `json:"billablePercentage"`
	BillableHours         float32 `json:"billableHours"`
	BillableHourlyRate    float32 `json:"billableHourlyRate"`
	CreatedAt             int64   `json:"createdAt"`
	UpdatedAt             int64   `json:"updatedAt"`
}

func (c *Claim) ConvertToDynamodbAttributes() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"claimId":             {S: aws.String(c.ClaimID)},
		"claimNumber":         {S: aws.String(c.ClaimNumber)},
		"insured":             {S: aws.String(c.Insured)},
		"claimCost":           {N: aws.String(fmt.Sprintf("%.2f", c.ClaimCost))},
		"closingCost":         {N: aws.String(fmt.Sprintf("%.2f", c.ClosingCost))},
		"status":              {S: aws.String(c.Status)},
		"description":         {S: aws.String(c.Description)},
		"clientId":            {S: aws.String(c.ClientID)},
		"assignedUserId":      {S: aws.String(c.AssignedUserID)},
		"creatorId":           {S: aws.String(c.CreatorID)},
		"workflowStep":        {S: aws.String(c.WorkflowStep)},
		"contactPerson":       {S: aws.String(c.ContactPerson)},
		"claimantAddress":     {S: aws.String(c.ClaimantAddress)},
		"dateOfLoss":          {N: aws.String(fmt.Sprintf("%d", c.DateOfLoss))},
		"claimantName":        {S: aws.String(c.ClaimantName)},
		"claimantContact":     {S: aws.String(c.ClaimantContact)},
		"claimantPhone":       {S: aws.String(c.ClaimantPhone)},
		"claimantEmail":       {S: aws.String(c.ClaimantEmail)},
		"contactPhone":        {S: aws.String(c.ContactPhone)},
		"contactEmail":        {S: aws.String(c.ContactEmail)},
		"clientBillingMethod": {S: aws.String(c.ClientBillingMethod)},
		"billablePercentage":  {N: aws.String(fmt.Sprintf("%.2f", c.BillablePercentage))},
		"billableHours":       {N: aws.String(fmt.Sprintf("%.2f", c.BillableHours))},
		"billableHourlyRate":  {N: aws.String(fmt.Sprintf("%.2f", c.BillableHourlyRate))},
		"createdAt":           {N: aws.String(fmt.Sprintf("%d", c.CreatedAt))},
		"updatedAt":           {N: aws.String(fmt.Sprintf("%d", c.UpdatedAt))},
	}
}
