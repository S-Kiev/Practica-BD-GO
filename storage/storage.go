package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"

	producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver del storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
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

func newMySQLDB() {
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

// FabricaProducto fabrica de producto.Storage
func FabricaProducto(driver Driver) (producto.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProducto(db), nil
	default:
		return nil, fmt.Errorf("Driver no implementado")
	}
}
