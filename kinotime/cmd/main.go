package main

import (
	"database/sql"
	"fmt"
	"os"

	"kinotime/internal/api"
	"kinotime/internal/configs"
	"kinotime/internal/repository"

	_ "github.com/lib/pq"
)

var (
	Db  *sql.DB
	Cfg *configs.Config
)

func main() {
	Cfg, err := configs.GetConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Db = repository.ConnectPostgre(Cfg.ConnPostgres)

	apiServer := api.NewServer(Cfg, Db)

	fallCount := 0
	err = apiServer.Start()
	if err != nil {
		fallCount++
		if fallCount < 3 {
			apiServer.Start()
		} else {
			os.Exit(1)
		}
	}
}
