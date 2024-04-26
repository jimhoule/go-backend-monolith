package objects

import (
	"app/gql"
)

var translationObject *gql.Object

func GetTranslationObject() *gql.Object {
	if translationObject == nil {
		translationObject = gql.CreateObject(gql.ObjectConfig{
			Name:        "Translation",
			Description: "Translation object",
			Fields: gql.Fields{
				"entityId": &gql.Field{
					Type: gql.ID,
					Description: "Associated entity ID",
				},
		
				"languageId": &gql.Field{
					Type: gql.ID,
					Description: "Associated language ID",
				},
		
				"text": &gql.Field{
					Type: gql.String,
					Description: "Text",
				},
				
				"type": &gql.Field{
					Type: gql.String,
					Description: "Type (TITLE, LABEL, DESCRIPTION)",
				},
			},
		})
	}

	return translationObject
}