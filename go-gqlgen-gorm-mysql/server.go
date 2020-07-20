package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leehom/go-gqlgen-gorm-mysql/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	"github.com/leehom/go-gqlgen-gorm-mysql/graph"
	"github.com/leehom/go-gqlgen-gorm-mysql/graph/generated"
)

const defaultPort = "8080"

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)
	db.Exec("CREATE DATABASE test_db")
	db.Exec("USE test_db")

	db.AutoMigrate(&model.Order{}, &model.Item{})
}

func main() {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
