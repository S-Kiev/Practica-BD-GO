package encabezadofactura

import (
	"database/sql"
	"time"
)

// Modelo del Encabezado de Factura
type Modelo struct {
	ID                 uint
	Cliente            string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

// Interfaz de almacenamiento que debe implementar un almacenamiento db
type Storage interface {
	Migrate() error
	CreateTransaction(*sql.Tx, *Modelo) error
}

// Servicio de Encabezado
type Servicio struct {
	storage Storage
}

// NewServicio retorna un puntero a de Servicio
func NewServicio(s Storage) *Servicio {
	return &Servicio{s}
}

// Migrate es usado para la migracion de producto
func (s *Servicio) Migrate() error {
	return s.storage.Migrate()
}
