package database

import (
	"database/sql"
	"final_project_backend/config"
	"fmt"
	"net/url"

	"github.com/lib/pq"
)

func NewDBConnection(conf *config.Database) (*sql.DB, error) {
	connURL := providePostgresURL(conf)

	parseURL, err := pq.ParseURL(connURL.String())
	if err != nil {
		return nil, fmt.Errorf("parsing postgres URI: %w", err)
	}

	conn, err := sql.Open("postgres", parseURL)
	if err != nil {
		return nil, fmt.Errorf("database open connection: %w", err)
	}

	return conn, nil
}

func providePostgresURL(conf *config.Database) *url.URL {
	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.Username, conf.Password),
		Path:   conf.Host + ":" + conf.Port + "/" + conf.Name,
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	return pgURL
}
