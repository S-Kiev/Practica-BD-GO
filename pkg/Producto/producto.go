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
	//Update(*Modelo) error
	//GetAll() (Modelos, error)
	//GetByID(uint) (*Modelo, error)
	//Delete(uint) error
}

// Servicio de Producto
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

// Create usado para crear un producto
func (s *Servicio) Create(m *Modelo) error {
	m.FechaCreacion = time.Now()
	return s.storage.Create(m)
}
