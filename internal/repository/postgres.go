package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testProject/config"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewPostgresPool(cfg *config.DBConf) (*sql.DB, error) {
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Printf("Error occured: %s", err)
		return nil, errors.New("database error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.TimeOut)*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB...")
	return db, nil
}
