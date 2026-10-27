package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedUrl := "https://www.wagslane.dev/index.xml"
	feed , err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Println("Here's your feed:")
	fmt.Printf("%+v", feed)
	return nil
}
