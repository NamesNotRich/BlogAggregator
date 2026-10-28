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
		return fmt.Errorf("user doesn't exist: %w", err)
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

	followArg := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), followArg)
	if err != nil {
		return fmt.Errorf("couldn't create follow: %w", err)
	}

	fmt.Println("Feed has been successfuly created:")
	PrintFeed(feed)
	fmt.Println("follow has been successfuly created!")
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list feeds")
	}

	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Printf(" %+v\n", feed)
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user doesn't exist: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed doesn't exist: %w", err)
	}

	arg := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("Feed followed successfuly!")
	fmt.Printf("* Feed Name: %v\n", feedFollow.FeedName)
	fmt.Printf("* Current User: %v\n", feedFollow.UserName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user doesn't exist: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	fmt.Printf("%s is following:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("- %s\n", feedFollow.FeedName)
	}
	return nil
}

func PrintFeed(feed database.Feed) {
	fmt.Printf(" * ID:            %v\n", feed.ID)
	fmt.Printf(" * Created:       %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:          %v\n", feed.Name)
	fmt.Printf(" * Url:           %v\n", feed.Url)
	fmt.Printf(" * UserID:        %v\n", feed.UserID)
}
