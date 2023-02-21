package itemfactura

import (
	//"database/sql"
	"time"
)

// Modelo del item de la factura
type Modelo struct {
	ID                  uint
	EncabezadoFacturaID uint
	ProductoID          uint
	FechaCreacion       time.Time
	FechaActualizacion  time.Time
}

// Modelos es un slice del Modelo (de los items de la factura)
type Modelos []*Modelo

// Interfaz de almacenamiento que debe implementar un almacenamiento db
type Storage interface {
	Migrate() error
	//	CreateTransaction(*sql.Tx, uint, Modelos) error
}

// Servicio de Item
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
