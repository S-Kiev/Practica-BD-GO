package storage

import (
	"database/sql"
	"fmt"
	//encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
)

const (
	MySQLMigrateEncabezadoFactura = `CREATE TABLE IF NOT EXISTS encabezado_factura(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		clinete VARCHAR(100) NOT NULL,
		fechaCreacion TIMESTAMP NOT NULL DEFAULT now(),
		fechaModificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	/*
		psqlCreateEncabezadoFactura = `INSERT INTO encabezado_factura(cliente) VALUES($1) RETURNING id, fechaCreacion`
	*/
)

// MySQLEncabezadoFactura es usado para trabajar con MySQL - encabezado
type MySQLEncabezadoFactura struct {
	db *sql.DB
}

// newMySQLEncabezadoFactura retorna un nuevo puntero a MySQLEncabezadoFactura
func NewMySQLEncabezadoFactura(db *sql.DB) *MySQLEncabezadoFactura {
	return &MySQLEncabezadoFactura{db}
}

func (p *MySQLEncabezadoFactura) Migrate() error {
	stmt, err := p.db.Prepare(MySQLMigrateEncabezadoFactura)
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
