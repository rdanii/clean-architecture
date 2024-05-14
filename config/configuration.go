package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Configuration adalah struct yang berisi slice dari Settings.
type Configuration struct {
	ENV []Settings `mapstructure:"env"` // Menggunakan tag mapstructure untuk memetakan field dari file konfigurasi.
}

// Settings adalah struct yang berisi nama dan nilai dari pengaturan.
type Settings struct {
	NAME  string // Nama dari variabel environment.
	VALUE string // Nilai dari variabel environment.
}

// InitConfig menginisialisasi konfigurasi berdasarkan environment yang diberikan.
func InitConfig(environment string) {
	// Mengatur tipe file konfigurasi yang akan dibaca sebagai YAML.
	viper.SetConfigType("yaml")

	// Memilih file konfigurasi berdasarkan environment.
	switch environment {
	case "local":
		// Mengatur nama file konfigurasi untuk environment "local".
		viper.SetConfigName("values.local")
	default:
		// Menampilkan pesan error dan menghentikan program jika environment tidak ditemukan.
		log.Fatalf("configuration for %s not found", environment)
	}

	// Menambahkan path untuk mencari file konfigurasi.
	viper.AddConfigPath("./config")
	// Mengatur viper untuk secara otomatis membaca variabel environment dari sistem operasi.
	viper.AutomaticEnv()

	// Membaca file konfigurasi.
	err := viper.ReadInConfig()
	if err != nil {
		// Menampilkan pesan error dan menghentikan program jika terjadi kesalahan saat membaca konfigurasi.
		log.Fatal(err)
	}

	// Membuat variabel env dari tipe Configuration.
	var env Configuration
	// Memetakan nilai yang dibaca oleh viper ke dalam struct env.
	viper.Unmarshal(&env)

	// Mengiterasi setiap nilai dalam slice ENV di struct Configuration.
	for _, value := range env.ENV {
		// Memeriksa apakah variabel environment dengan nama yang diberikan sudah ada.
		if _, ok := os.LookupEnv(value.NAME); !ok {
			// Jika variabel environment belum ada, mengaturnya dengan nilai yang diberikan.
			os.Setenv(value.NAME, value.VALUE)
		}
	}
}
