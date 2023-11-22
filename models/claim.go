package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Claim struct {
	ClaimID               string `json:"claimId"`
	ClaimNumber           string `json:"claimNumber"`
	Insured               string `json:"insured"`
	ClaimCost             string `json:"claimCost"`
	ClosingCost           string `json:"closingCost"`
	Status                string `json:"status"`
	Description           string `json:"description"`
	ClientID              string `json:"clientId"`
	AssignedUserID        string `json:"assignedUserId"`
	AssignedUserFirstName string `json:"assignedUserFirstName"`
	AssignedUserLastName  string `json:"assignedUserLastName"`
	CreatorID             string `json:"creatorId"`
	CreatorFirstName      string `json:"creatorFirstName"`
	CreatorLastName       string `json:"creatorLastName"`
	WorkflowStep          string `json:"workflowStep"`
	ContactPerson         string `json:"contactPerson"`
	ClaimantAddress       string `json:"claimantAddress"`
	DateOfLoss            string `json:"dateOfLoss"`
	ClaimantName          string `json:"claimantName"`
	ClaimantContact       string `json:"claimantContact"`
	ClaimantPhone         string `json:"claimantPhone"`
	ClaimantEmail         string `json:"claimantEmail"`
	ContactPhone          string `json:"contactPhone"`
	ContactEmail          string `json:"contactEmail"`
	ClientBillingMethod   string `json:"clientBillingMethod"`
	BillablePercentage    string `json:"billablePercentage"`
	BillableHours         string `json:"billableHours"`
	BillableHourlyRate    string `json:"billableHourlyRate"`
	CreatedAt             string `json:"createdAt"`
	UpdatedAt             string `json:"updatedAt"`
}

func (c *Claim) ConvertToDynamodbAttributes() (map[string]types.AttributeValue, error) {
	// Manually format float fields as currency strings
	// claimCost := fmt.Sprintf("%.2f", c.ClaimCost)
	// closingCost := fmt.Sprintf("%.2f", c.ClosingCost)
	// billablePercentage := fmt.Sprintf("%.2f", c.BillablePercentage)
	// billableHours := fmt.Sprintf("%.2f", c.BillableHours)
	// billableHourlyRate := fmt.Sprintf("%.2f", c.BillableHourlyRate)

	// Create a temporary map with formatted strings for float fields
	tempClaimMap := map[string]interface{}{
		"claimId":             c.ClaimID,
		"claimNumber":         c.ClaimNumber,
		"insured":             c.Insured,
		"claimCost":           c.ClaimCost,
		"closingCost":         c.ClosingCost,
		"status":              c.Status,
		"description":         c.Description,
		"clientId":            c.ClientID,
		"assignedUserId":      c.AssignedUserID,
		"creatorId":           c.CreatorID,
		"workflowStep":        c.WorkflowStep,
		"contactPerson":       c.ContactPerson,
		"claimantAddress":     c.ClaimantAddress,
		"dateOfLoss":          c.DateOfLoss,
		"claimantName":        c.ClaimantName,
		"claimantContact":     c.ClaimantContact,
		"claimantPhone":       c.ClaimantPhone,
		"claimantEmail":       c.ClaimantEmail,
		"contactPhone":        c.ContactPhone,
		"contactEmail":        c.ContactEmail,
		"clientBillingMethod": c.ClientBillingMethod,
		"billablePercentage":  c.BillablePercentage,
		"billableHours":       c.BillableHours,
		"billableHourlyRate":  c.BillableHourlyRate,
		"createdAt":           c.CreatedAt,
		"updatedAt":           c.UpdatedAt,
	}

	av, err := attributevalue.MarshalMap(tempClaimMap)
	if err != nil {
		return nil, err
	}

	return av, nil
}
