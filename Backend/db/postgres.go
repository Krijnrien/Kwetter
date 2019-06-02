package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/krijnrien/Kwetter/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertKweet(ctx context.Context, Kweet schema.Kweet) error {
	_, err := r.db.Exec("INSERT INTO Kweets(id, body, created_at) VALUES($1, $2, $3)", Kweet.ID, Kweet.Body, Kweet.CreatedAt)
	return err
}

func (r *PostgresRepository) ListKweets(ctx context.Context, skip uint64, take uint64) ([]schema.Kweet, error) {
	rows, err := r.db.Query("SELECT * FROM Kweets ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Kweets
	Kweets := []schema.Kweet{}
	for rows.Next() {
		Kweet := schema.Kweet{}
		if err = rows.Scan(&Kweet.ID, &Kweet.Body, &Kweet.CreatedAt); err == nil {
			Kweets = append(Kweets, Kweet)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Kweets, nil
}
