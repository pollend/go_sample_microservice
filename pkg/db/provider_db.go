package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	_ "github.com/lib/pq"
)

type ProviderDB sqlx.DB

func NewProviderDatabase() *ProviderDB {
	for {
		db, err := sqlx.Connect("postgres", "postgres://petstore_api:password@pg:5432/petstore_api?sslmode=disable")
		if err == nil {
			return (*ProviderDB)(db)
		}
		fmt.Println(err.Error())
		fmt.Println("Failed to establish connection trying in 5 seconds")
		time.Sleep(time.Second * 5)
	}
}
