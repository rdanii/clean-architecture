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

	// Mendefinisikan flag untuk menentukan environment, default "local"
	envF := flag.String("env", "local", "define environment")
	// Mem-parsing flag yang diberikan dari command line
	flag.Parse()

	// Mengambil nilai dari flag environment
	env := *envF
	// Menginisialisasi konfigurasi berdasarkan environment yang diberikan
	config.InitConfig(env)

	// Menghubungkan ke PostgreSQL dan menangani error jika terjadi
	postgresConn, errPostgres := config.ConnectPostgreSQL()
	if errPostgres != nil {
		log.Panic("error PostgreSQL connection: ", errPostgres)
	}
	// Menutup koneksi PostgreSQL saat program selesai
	defer func() {
		if sqlDB, err := postgresConn.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	// Menginisialisasi user dari database menggunakan GORM v2
	errAuth := auth.InitializeUserFromDBGormV2(postgresConn)
	if errAuth != nil {
		log.Println(errAuth.Error())
	}

	// Menginisialisasi router aplikasi
	router := app.InitRouter(postgresConn)

	log.Println("routes Initialized")

	// Mendapatkan port server dari konfigurasi
	port := viper.GetString("PORT")
	// Menginisialisasi server HTTP dengan router yang sudah diatur
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server Initialized, listening at port: ", port)

	// Menjalankan server HTTP dalam goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Membuat channel untuk menangani sinyal sistem (seperti interrupt)
	quit := make(chan os.Signal, 1)
	// Mengatur channel untuk menerima sinyal interrupt
	signal.Notify(quit, os.Interrupt)
	// Menunggu sinyal interrupt
	<-quit
	log.Println("Shutdown Server ...")

	// Membuat konteks dengan timeout 5 detik untuk shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Menutup server dengan aman dalam konteks yang diberikan
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
