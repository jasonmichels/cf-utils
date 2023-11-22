package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Note struct {
	NoteID           string `json:"noteId"`
	ClaimID          string `json:"claimId"`
	Body             string `json:"body"`
	CreatorID        string `json:"creatorId"`
	CreatorFirstName string `json:"creatorFirstName"`
	CreatorLastName  string `json:"creatorLastName"`
	CreatedAt        int64  `json:"createdAt"`
	UpdatedAt        int64  `json:"updatedAt"`
}

func (n *Note) ConvertToDynamodbAttributes() (map[string]types.AttributeValue, error) {
	tempNoteMap := map[string]interface{}{
		"noteId":    n.NoteID,
		"claimId":   n.ClaimID,
		"body":      n.Body,
		"creatorId": n.CreatorID,
		"createdAt": n.CreatedAt,
		"updatedAt": n.UpdatedAt,
	}

	av, err := attributevalue.MarshalMap(tempNoteMap)
	if err != nil {
		return nil, err
	}

	return av, nil
}

func EnrichNotes(notes []Note, users []User) ([]Note, error) {
	enrichedNotes := make([]Note, len(notes))

	for i, note := range notes {
		enrichedNote := note

		// Fetch and set creator user details
		if creatorUser, err := GetUserByID(note.CreatorID, users); err == nil {
			enrichedNote.CreatorFirstName = creatorUser.FirstName
			enrichedNote.CreatorLastName = creatorUser.LastName
		}

		enrichedNotes[i] = enrichedNote
	}

	return enrichedNotes, nil
}
