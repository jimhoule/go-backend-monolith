package gql

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var ID = graphql.ID
var Int = graphql.Int
var String = graphql.String

type Object = graphql.Object
type ObjectConfig = graphql.ObjectConfig
type Fields = graphql.Fields
type Field = graphql.Field
type FieldConfigArgument = graphql.FieldConfigArgument
type ArgumentConfig = graphql.ArgumentConfig
type List = graphql.List
type ResolveParams = graphql.ResolveParams

func CreateObject(objectConfig ObjectConfig) *Object {
	return graphql.NewObject(objectConfig)
}

func CreateList(objectType graphql.Type) *List {
	return graphql.NewList(objectType)
}

type Server struct{
	RootQueryFields Fields
}

// Adds query or mutation
func (s *Server) Add(name string, field *Field) {
	s.RootQueryFields[name] = field
}

func (s *Server) ServeGQL() *handler.Handler {
	rootQuery := ObjectConfig{
		Name:   "RootQuery",
		Fields: s.RootQueryFields,
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: CreateObject(rootQuery),
	})
	if err != nil {
		fmt.Printf("Failed to create graphql schema: %v", err)
	}
	
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})
}

func (*Server) ServeSandbox() http.HandlerFunc {
	sandboxHtml := []byte(`
		<!DOCTYPE html>
		<html lang="en">
			<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
				<div id="sandbox" style="height:100vh; width:100vw;"></div>

				<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
				<script>
					new window.EmbeddedSandbox({
						target: "#sandbox",
						// Pass through your server href if you are embedding on an endpoint.
						// Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
						initialEndpoint: "http://localhost:3000/graphql",
					});
					
					// advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
				</script>
			</body>
		</html>
	`)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(sandboxHtml)
	})
}

var gqlServer *Server

func Get() *Server {
	if gqlServer == nil {
		gqlServer = &Server{
			RootQueryFields: Fields{},
		}
	}

	return gqlServer
}