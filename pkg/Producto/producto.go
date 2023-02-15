package producto

import "time"

//Modelo de Producto
type Modelo struct {
	ID                 uint
	Nombre             string
	Detalle            string
	Precio             int
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}

// Slice de Modelo
type Modelos []*Modelo

// Interfaz de almacenamiento que debe implementar un almacenamiento db
type Storage interface {
	Migrate() error
	Create(*Modelo) error
	Update(*Modelo) error
	GetAll() (Modelos, error)
	GetByID(uint) (*Modelo, error)
	Delete(uint) error
}
