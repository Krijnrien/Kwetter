package db

import (
	"context"

	"github.com/krijnrien/Kwetter/schema"
)

type Repository interface {
	Close()
	InsertKweet(ctx context.Context, Kweet schema.Kweet) error
	ListKweets(ctx context.Context, skip uint64, take uint64) ([]schema.Kweet, error)
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

func ListKweets(ctx context.Context, skip uint64, take uint64) ([]schema.Kweet, error) {
	return impl.ListKweets(ctx, skip, take)
}
