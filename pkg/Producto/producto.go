package producto

import (
	//"errors"
	"fmt"
	"strings"
	"time"
)

// Modelo de Producto
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

func (m *Modelo) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Nombre, m.Detalle, m.Precio,
		m.FechaCreacion.Format("2006-01-02"), m.FechaActualizacion.Format("2006-01-02"))
}

func (m Modelos) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"id", "nombre", "detalle", "precio", "fechaCreacion", "fechaModificacion"))
	for _, modelo := range m {
		builder.WriteString(modelo.String() + "\n")
	}
	return builder.String()
}

// Interfaz de almacenamiento que debe implementar un almacenamiento db
type Storage interface {
	Migrate() error
	Create(*Modelo) error
	//Update(*Modelo) error
	GetAll() (Modelos, error)
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

// GetAll usado para obtener todos los productos
func (s *Servicio) GetAll() (Modelos, error) {
	return s.storage.GetAll()
}

// GetByID es usado para obtener un producto especifico
//func (s *Servicio) GetByID(id uint) (*Modelo, error) {
//	return s.storage.GetByID(id)
//}

//var (
//	ErrIDNotFound = errors.New("El producto no contiene un ID")
//)

// Update usado para actualizar un producto
//func (s *Servicio) Update(m *Modelo) error {
//	if m.ID == 0 {
//		return ErrIDNotFound
//	}
//
//	m.FechaActualizacion = time.Now()
//
//	return s.storage.Update(m)
//}

// Delete para eliminar un producto
//func (s *Servicio) Delete(id uint) error {
//	return s.storage.Delete(id)
//}
