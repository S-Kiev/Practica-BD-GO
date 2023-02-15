package encabezadofactura

import "time"

// Modelo del Encabezado de Factura
type Modelo struct {
	ID                 uint
	Cliente            string
	FechaCreacion      time.Time
	FechaActualizacion time.Time
}
