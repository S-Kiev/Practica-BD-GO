package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS productos(
		id SERIAL NOT NULL,
		nombre VARCHAR(25) NOT NULL,
		detalle VARCHAR(100),
		precio INT NOT NULL,
		fechaCreacion TIMESTAMP NOT NULL DEFAULT now(),
		fechaModificacion TIMESTAMP,
		CONSTRAINT productos_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateProduct = `INSERT INTO productos(nombre, detalle, precio, fechaCreacion) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct = `SELECT id, nombre, detalle, precio, 
	fechaCreacion, fechaModificacion
	FROM productos`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE productos SET nombre = $1, detalle = $2,
	precio = $3, fechaModificacion = $4 WHERE id = $5`
	psqlDeleteProduct = `DELETE FROM productos WHERE id = $1`
)

// psqlProducto usado para trabajar con Postgress - producto
type PsqlProduct struct {
	db *sql.DB
}

// newPsqlProduct retorna un nuevo puntero a PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
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
