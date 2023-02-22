package storage

import (
	"database/sql"
	"fmt"

	itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
)

const (
	MySQLMigrateItemFactura = `CREATE TABLE IF NOT EXISTS item_factura(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		encabezado_factura_id INT NOT NULL,
		producto_id INT NOT NULL,
		fechaCreacion TIMESTAMP NOT NULL DEFAULT now(),
		fechaModificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT item_factura_encabezado_factura_id_fk FOREIGN KEY (encabezado_factura_id) REFERENCES encabezado_factura (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT item_factura_producto_id_fk FOREIGN KEY (producto_id) REFERENCES productos (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`

	mysqlCreateItemFactura = `INSERT INTO item_factura(encabezado_factura_id, producto_id) VALUES(?, ?)`
)

// MySQLItemFactura usado para trabajar con MySQL - item de factura
type MySQLItemFactura struct {
	db *sql.DB
}

// newMySQLItemFactura retorna un nuevo puntero a MySQLItemFactura
func NewMySQLItemFactura(db *sql.DB) *MySQLItemFactura {
	return &MySQLItemFactura{db}
}

func (p *MySQLItemFactura) Migrate() error {
	stmt, err := p.db.Prepare(MySQLMigrateItemFactura)
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
func (p *MySQLItemFactura) CreateTransaction(tx *sql.Tx, encabezadoID uint, ms itemfactura.Modelos) error {
	stmt, err := tx.Prepare(mysqlCreateItemFactura)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {

		resultado, err := stmt.Exec(encabezadoID, item.ProductoID)
		if err != nil {
			return err
		}

		id, err := resultado.LastInsertId()
		if err != nil {
			return err
		}

		item.ID = uint(id)

	}

	return nil
}
