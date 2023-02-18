package storage

import (
	"database/sql"
	"fmt"

	encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
	factura "github.com/S-Kiev/Practica-BD-GO/pkg/Factura"
	itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
)

// PsqlFactura usado para trabajar con postgres - factura
type PsqlFactura struct {
	db           *sql.DB
	encabezado   encabezadofactura.Storage
	itemsFactura itemfactura.Storage
}

// NewPsqlFactura returna un nuevo puntero de PsqlFactura
func NewPsqlFactura(db *sql.DB, encabezado encabezadofactura.Storage, item itemfactura.Storage) *PsqlFactura {
	return &PsqlFactura{
		db:           db,
		encabezado:   encabezado,
		itemsFactura: item,
	}
}

// Create implementa la interface factura.Storage
func (f *PsqlFactura) Create(m *factura.Modelo) error {
	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	if err := f.encabezado.CreateTransaction(tx, m.Encabezado); err != nil {
		tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Encabezado.ID)

	if err := f.itemsFactura.CreateTransaction(tx, m.Encabezado.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("Items: %w", err)
	}
	fmt.Printf("items creados: %d \n", len(m.Items))

	return tx.Commit()
}
