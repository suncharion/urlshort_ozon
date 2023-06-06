package storagepostgres

import (
	"context"
	"fmt"
	"os"
	"ozon/internal/model"

	"github.com/jackc/pgx/v5"
)

type StoragePostgres struct {
	data map[string]*model.ShortenedUrl
	conn *pgx.Conn
}

func NewStoragePostgres() (*StoragePostgres, error) {

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:qwe@localhost:5432/urls")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &StoragePostgres{
		data: make(map[string]*model.ShortenedUrl),
		conn: conn,
	}, nil
}

func (sm *StoragePostgres) Put(url *model.ShortenedUrl) error {
	fmt.Println(url)
	row := sm.conn.QueryRow(context.Background(), "INSERT INTO urls (short, original) VALUES ($1, $2) RETURNING short", url.Short, url.Original)
	short := ""
	err := row.Scan(&short)
	if err != nil {
		return err
	}

	return nil
}

func (sm *StoragePostgres) Get(shortUrl string) (*model.ShortenedUrl, error) {
	var short, original string
	err := sm.conn.QueryRow(context.Background(), "SELECT short, original FROM urls WHERE short = $1", shortUrl).Scan(&short, &original)
	if err != nil {
		return nil, fmt.Errorf("Not found")
	}
	return &model.ShortenedUrl{
		Short:    short,
		Original: original,
	}, nil
}
