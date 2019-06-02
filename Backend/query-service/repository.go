package main

import (
	"context"

	"github.com/krijnrien/Kwetter/schema"
)

type Repository interface {
	Close()
	InsertKweet(ctx context.Context, Kweet schema.Kweet) error
	SearchKweets(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Kweet, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertKweet(ctx context.Context, Kweet schema.Kweet) error {
	return impl.InsertKweet(ctx, Kweet)
}

func SearchKweets(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Kweet, error) {
	return impl.SearchKweets(ctx, query, skip, take)
}
