package storage

import (
	"database/sql"
	"fmt"
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
	psqlCreateItemFactura = `INSERT INTO item_factura(cliente) VALUES($1) RETURNING id, fechaCreacion`
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
