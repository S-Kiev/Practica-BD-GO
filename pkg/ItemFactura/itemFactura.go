package itemfactura

import "time"

// Modelo del item de la factura
type Model struct {
	ID                  uint
	EncabezadoFacturaID uint
	ProductoID          uint
	FechaCreacion       time.Time
	FechaActualizacion  time.Time
}
