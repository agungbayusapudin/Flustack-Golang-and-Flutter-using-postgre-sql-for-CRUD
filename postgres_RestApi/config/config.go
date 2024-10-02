package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {
	// load env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("ini eror", err)
	}

	// connection to database
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatal(err)
	}

	// check database
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sukses Konek ke Db!")

	return db

}

// marshal digunakan untuk mengubah struktur data menjadi JSON

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
