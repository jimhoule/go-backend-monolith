package objects

import (
	"app/gql"
	"app/translations/presenters/gql/objects"
)

var movieObject *gql.Object

func GetMovieObject() *gql.Object {
	if movieObject == nil {
		movieObject = gql.CreateObject(gql.ObjectConfig{
			Name:        "Movie",
			Description: "Movie object",
			Fields: gql.Fields{
				"id": &gql.Field{
					Type: gql.ID,
					Description: "Unique identifier",
				},

				"titles": &gql.Field{
					Type: gql.CreateList(objects.GetTranslationObject()),
					Description: "Titles translations",
				},

				"descriptions": &gql.Field{
					Type: gql.CreateList(objects.GetTranslationObject()),
					Description: "Descriptions translations",
				},

				"genreId": &gql.Field{
					Type: gql.ID,
					Description: "Associated genre ID",
				},
			},
		})
	}

	return movieObject
}