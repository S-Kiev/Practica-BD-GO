package factura

import (
	encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
	itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
)

// Modelo de Factura
type Modelo struct {
	Encabezado *encabezadofactura.Modelo
	Items      itemfactura.Modelos
}

// Storage interface que debe implementar el almaceniento de BD
type Storage interface {
	Create(*Modelo) error
}

// Servicio de factura
type Servicio struct {
	storage Storage
}

// NewService returna un puntero de Servicio
func NewService(s Storage) *Servicio {
	return &Servicio{s}
}

// Crea una nueva factura
func (s *Servicio) Create(m *Modelo) error {
	return s.storage.Create(m)
}
