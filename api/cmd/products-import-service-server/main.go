package main

import (
	"database/sql"
	"log"

	"github.com/TamerB/products-import-service/api"
	"github.com/TamerB/products-import-service/config"

	db "github.com/TamerB/products-import-service/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	config := config.NewConfig()
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	db.DBStore = db.NewStore(conn).(*db.SQLStore)

	server := api.NewServer()

	err = server.Start(config.URL)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
