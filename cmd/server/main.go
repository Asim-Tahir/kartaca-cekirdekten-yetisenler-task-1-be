package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
	"github.com/shellbear/go-graphql-example/api"
	"github.com/shellbear/go-graphql-example/configs"
	
	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configs.InitDatabase()
	defer configs.DB.Close()

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}})))

	log.Printf("🚀 GraphQL Playground Server running on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
