package main

import (
	"graphql1/database"
	"graphql1/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	ser := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.POST("/query", gin.WrapH(ser))
	router.GET("/", gin.WrapH(playground.Handler("grap", "/query")))
	router.Run()
}
