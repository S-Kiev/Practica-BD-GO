package itemfactura

import "time"

// Modelo del item de la factura
type Modelo struct {
	ID                  uint
	EncabezadoFacturaID uint
	ProductoID          uint
	FechaCreacion       time.Time
	FechaActualizacion  time.Time
}

// Interfaz de almacenamiento que debe implementar un almacenamiento db
type Storage interface {
	Migrate() error
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
