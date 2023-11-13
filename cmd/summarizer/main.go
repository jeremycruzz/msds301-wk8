package main

import (
	"flag"
	"log"

	"github.com/jeremycruzz/msds301-wk8/internal/controller"
	"github.com/jeremycruzz/msds301-wk8/internal/server"
	"github.com/jeremycruzz/msds301-wk8/pkg/chatgpt"
	"github.com/jeremycruzz/msds301-wk8/pkg/sqlitedb"
	"github.com/jeremycruzz/msds301-wk8/pkg/summarizer"
	"github.com/jeremycruzz/msds301-wk8/pkg/wikiapi"
)

func main() {
	apiKey := flag.String("apikey", "", "API key for chatgpt")
	flag.Parse()

	if *apiKey == "" {
		log.Fatal("API key is required. Start the server with -apikey flag.")
	}

	port := flag.String("p", ":8080", "Port for server to listen on")

	wikiApi := wikiapi.New()
	chatGpt := chatgpt.New(*apiKey)
	db := sqlitedb.NewSqlite("./data/summaries.sqlite")

	summarizerService := summarizer.New(wikiApi, chatGpt, db)
	controller := controller.NewSummarizerController(summarizerService)
	router := server.NewRouter(controller)
	server.Start(router, *port)
}
