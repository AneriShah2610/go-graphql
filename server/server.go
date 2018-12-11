package main

import (
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
	graphql "github.com/aneri/go-graphql-gqlgen"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := graphql.Config{Resolvers: &graphql.Resolver{}}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(c)))
	//http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(c), handler.ComplexityLimit(5))) // To limit complexity to 5
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
