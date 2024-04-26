package resolvers

import (
	"app/gql"
	"app/movies/application/services"
	"app/movies/presenters/gql/objects"
	"fmt"
)

type MoviesResolver struct {
	MoviesService *services.MoviesService
}

func (mr *MoviesResolver) GetFindAllMoviesQuery() *gql.Field {
	return &gql.Field{
		Name:        "FindAllMovies",
		Type:        gql.CreateList(objects.GetMovieObject()),
		Description: "Finds all movies",
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			return mr.MoviesService.FindAll()
		},
	}
}

func (mr *MoviesResolver) GetFindMovieByIdQuery() *gql.Field {
	return &gql.Field{
		Name: "FindMovieById",
		Type: objects.GetMovieObject(),
		Description: "Finds a movie by ID",
		Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{
				Type: gql.ID,
				Description: "Movie ID",
			},
		},
		Resolve: func(params gql.ResolveParams) (interface{}, error) {
			id, isValid := params.Args["id"]
			if !isValid {
				return nil, fmt.Errorf("movie id %s is invalid", id)
			}

			return mr.MoviesService.FindById(id.(string))
		},
	}
}