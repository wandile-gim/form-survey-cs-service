package repository

import (
	"context"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"fmt"
	"form-survey-cs-service/internal/config"
	"form-survey-cs-service/internal/repository/ent"
	"form-survey-cs-service/internal/repository/ent/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

func pgxConnection() *pgxpool.Pool {
	d := config.DatabaseConfiguration
	address := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", d.Username, d.Password, d.Host, d.Port, d.UseDatabase)
	poolConfig, err := pgxpool.ParseConfig(address)
	poolConfig.MaxConns = 100
	poolConfig.MinConns = 0
	poolConfig.MaxConnLifetime = time.Minute * 2
	pool, err := pgxpool.NewWithConfig(context.TODO(), poolConfig)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}

func Open() *ent.Client {
	pool := pgxConnection()
	db := stdlib.OpenDBFromPool(pool)
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv), ent.Debug())
}

func Migration() {
	initCreateErr := config.DatabaseConfiguration.CreateDatabaseIfNotExists("postgres")
	if initCreateErr != nil {
		log.Fatalf("initCreateErr %v", initCreateErr.Error())
	}

	d := config.DatabaseConfiguration
	address := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", d.Username, d.Password, d.Host, d.Port, d.UseDatabase)

	// Run migration.
	ctx := context.Background()
	client, err := ent.Open("postgres", address)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	defer client.Close()

	if cErr := client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); cErr != nil {
		log.Fatalf("failed creating schema resources: %v", cErr)
	}

}
