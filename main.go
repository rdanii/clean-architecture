package main

import (
	"clean-architecture/app"
	"clean-architecture/auth"
	"clean-architecture/config"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
)

func main() {

	envF := flag.String("env", "local", "define environment")
	flag.Parse()

	env := *envF
	config.InitConfig(env)

	postgresConn, errPostgres := config.ConnectPostgreSQL()
	if errPostgres != nil {
		log.Panic("error PostgreSQL connection: ", errPostgres)
	}
	defer func() {
		if sqlDB, err := postgresConn.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	errAUth := auth.InitializeUserFromDBGormV2(postgresConn)
	if errAUth != nil {
		log.Println(errAUth.Error())
	}

	router := app.InitRouter(postgresConn)

	log.Println("routes Initialized")

	port := viper.GetString("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server Initialized, listening at port: ", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
