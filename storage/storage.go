package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://nombreUsuario:clave@localhost:7530/nombreBD?sslmode=disable")
		if err != nil {
			log.Fatalf("No se pudo abrir la BD: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No se puedo hacer ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// Pool retorna una unica instacia de db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// Apartir de aqui sera para MySQL

func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "S-Kiev:sakura1997@tcp(localhost:3306)/bd-cursogo?parseTime=true")
		if err != nil {
			log.Fatalf("No se pudo conectar con la Base de Datos: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer ping: %v", err)
		}

		fmt.Println("conectado a MySQL")
	})
}
