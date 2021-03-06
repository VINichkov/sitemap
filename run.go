package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sitemap/creator"
	"sitemap/environment"
	"time"
)


func main()  {
	// logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
			TimeFormat: time.RFC3339,
		},
	)

	// DB
	environment.LoadEnviroment("config/settings.json")
	var db_name string
	var connectionString string
	if os.Getenv("DATABASE_URL") != "" {
		connectionString = os.Getenv("DATABASE_URL")
	} else {
		if os.Getenv("DB_NAME") == "" {
			db_name = fmt.Sprintf("mango_%s", os.Getenv("APPLICATION_ENV"))
		} else {
			db_name = fmt.Sprintf("mango_%s", os.Getenv("DB_NAME"))
		}

		connectionString = fmt.Sprintf("host=localhost password=%s user=%s dbname=%s sslmode=disable",
			os.Getenv("MONGO_DATABASE_PASSWORD"),
			os.Getenv("MONGO_DATABASE_USER"),
			db_name)
	}
	//log.Info().Msg(connectionString)
	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}else{
		log.Info().Msg("Connected to DB!")
	}

	log.Info().Msg("Start")
	service := creator.NewCreator(db)
	service.Start()

	defer db.Close()

}
