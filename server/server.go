package main

import (
	middleware2 "github.com/graphql-stack/backend-go/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-stack/backend-go"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(cors.Default().Handler)

	router.Use(middleware2.AuthMiddleware())
	router.Use(middleware2.DataloadersMiddleware())
	c := backend_go.Config{Resolvers: &backend_go.Resolver{}}
	gql := handler.GraphQL(backend_go.NewExecutableSchema(c))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", gql)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
