package storage

import (
	"database/sql"
	"fmt"

	itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
)

const (
	psqlMigrateItemFactura = `CREATE TABLE IF NOT EXISTS item_factura(
		id SERIAL NOT NULL,
		encabezado_factura_id INT NOT NULL,
		producto_id INT NOT NULL,
		fechaCreacion TIMESTAMP NOT NULL DEFAULT now(),
		fechaModificacion TIMESTAMP,
		CONSTRAINT item_factura_id_pk PRIMARY KEY (id),
		CONSTRAINT item_factura_encabezado_factura_id_fk FOREIGN KEY (encabezado_factura_id) REFERENCES encabezado_factura (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT item_factura_producto_id_fk FOREIGN KEY (producto_id) REFERENCES productos (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateItemFactura = `INSERT INTO item_factura(encabezado_factura_id, producto_id) VALUES($1, $2) RETURNING id, fechaCreacion`
)

// PsqlItemFactura usado para trabajar con Postgress - item de factura
type PsqlItemFactura struct {
	db *sql.DB
}

// newPsqlItemFactura retorna un nuevo puntero a PsqlItemFactura
func NewPsqlItemFactura(db *sql.DB) *PsqlItemFactura {
	return &PsqlItemFactura{db}
}

func (p *PsqlItemFactura) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de item de factura ejecutada correctamente")
	return nil
}

// CreateTransaction implementa la interface itemFactura.Storage
func (p *PsqlItemFactura) CreateTransaction(tx *sql.Tx, encabezadoID uint, ms itemfactura.Modelos) error {
	stmt, err := tx.Prepare(psqlCreateItemFactura)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(encabezadoID, item.ProductoID).Scan(
			&item.ID,
			&item.FechaCreacion,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
