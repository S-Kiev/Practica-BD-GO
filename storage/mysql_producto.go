package storage

import (
	"database/sql"
	"fmt"
	//producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
)

const (
	MySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS productos(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		nombre VARCHAR(25) NOT NULL,
		detalle VARCHAR(100),
		precio INT NOT NULL,
		fechaCreacion TIMESTAMP NOT NULL DEFAULT now(),
		fechaModificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	/*
		psqlCreateProduct = `INSERT INTO productos(nombre, detalle, precio, fechaCreacion) VALUES($1, $2, $3, $4) RETURNING id`
		psqlGetAllProduct = `SELECT id, nombre, detalle, precio,
		fechaCreacion, fechaModificacion
		FROM productos`
		psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
		psqlUpdateProduct  = `UPDATE productos SET nombre = $1, detalle = $2,
		precio = $3, fechaModificacion = $4 WHERE id = $5`
		psqlDeleteProduct = `DELETE FROM productos WHERE id = $1`
	*/
)

// MySQLProductoo usado para trabajar con MySQL - producto
type MySQLProducto struct {
	db *sql.DB
}

// newMySQLProducto retorna un nuevo puntero a MySQLProducto
func NewMySQLProducto(db *sql.DB) *MySQLProducto {
	return &MySQLProducto{db}
}

func (p *MySQLProducto) Migrate() error {
	stmt, err := p.db.Prepare(MySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}
