package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shihanng/devto/pkg/devto"
)

func main() {
	auth := context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: os.Getenv("DEV_TOKEN"),
	})

	client := devto.NewAPIClient(devto.NewConfiguration())
	articles, _, err := client.ArticlesApi.GetUserAllArticles(auth, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range articles {
		fmt.Println(a.Title, a.Id)
	}
}
