package main

import (
	"context"
	"fmt"
	"time"

	"github.com/NamesNotRich/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Errorf("user doesn't exist: %w", err)
	}

	arg := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed has been successfuly created:")
	PrintFeed(feed)
	return nil
}

func PrintFeed(feed database.Feed) {
	fmt.Printf(" * ID:         %v\n", feed.ID)
	fmt.Printf(" * Name:       %v\n", feed.Name)
	fmt.Printf(" * Url:        %v\n", feed.Url)
	fmt.Printf(" * UserID:     %v\n", feed.UserID)
}
