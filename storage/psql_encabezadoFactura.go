package storage

import (
	"database/sql"
	"fmt"

	encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
)

const (
	psqlMigrateEncabezadoFactura = `CREATE TABLE IF NOT EXISTS encabezado_factura(
		id SERIAL NOT NULL,
		clinete VARCHAR(100) NOT NULL,
		fechaCreacion TIMESTAMP NOT NULL DEFAULT now(),
		fechaModificacion TIMESTAMP,
		CONSTRAINT encabezadoFactura_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateEncabezadoFactura = `INSERT INTO encabezado_factura(cliente) VALUES($1) RETURNING id, fechaCreacion`
)

// PsqlEncabezadoFacturao usado para trabajar con Postgress - encabezado
type PsqlEncabezadoFactura struct {
	db *sql.DB
}

// newPsqlEncabezadoFactura retorna un nuevo puntero a PsqlEncabezadoFactura
func NewPsqlEncabezadoFactura(db *sql.DB) *PsqlEncabezadoFactura {
	return &PsqlEncabezadoFactura{db}
}

func (p *PsqlEncabezadoFactura) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de encabezado de factura ejecutada correctamente")
	return nil
}

// CreateTx implementa la interface encabezadoFactura.Storage
func (p *PsqlEncabezadoFactura) CreateTransaction(tx *sql.Tx, encabezado *encabezadofactura.Modelo) error {
	stmt, err := tx.Prepare(psqlCreateEncabezadoFactura)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(encabezado.Cliente).Scan(&encabezado.ID, &encabezado.FechaCreacion)
}
