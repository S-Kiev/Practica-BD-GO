package storage

import (
	"database/sql"
	"fmt"

	producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
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
type psqlProduct struct {
	db *sql.DB
}

// newpsqlProduct retorna un nuevo puntero a psqlProduct
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

func (p *psqlProduct) Migrate() error {
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

// Create implementa la interface Producto.Storage
func (p *psqlProduct) Create(m *producto.Modelo) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Nombre,
		stringToNull(m.Detalle),
		m.Precio,
		m.FechaCreacion,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("se creo el producto correctamente")
	return nil
}

// GetAll implementa la interface product.Storage
func (p *psqlProduct) GetAll() (producto.Modelos, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(producto.Modelos, 0)
	for rows.Next() {

		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetByID implementa la interface producto.Storage
func (p *psqlProduct) GetByID(id uint) (*producto.Modelo, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &producto.Modelo{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

func scanRowProduct(s scanner) (*producto.Modelo, error) {
	m := &producto.Modelo{}

	//Si hay segistros que pueden venir nulos de la BD
	detalleNull := sql.NullString{}
	fechaActualizacionNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Nombre,
		&m.Precio,
		&detalleNull.String,
		&m.FechaCreacion,
		&fechaActualizacionNull,
	)
	if err != nil {
		return &producto.Modelo{}, err
	}

	if detalleNull.Valid { // Verifica si el valor es válido antes de asignarlo
		m.Detalle = detalleNull.String
	}

	if fechaActualizacionNull.Valid { // Verifica si el valor es válido antes de asignarlo
		m.FechaActualizacion = fechaActualizacionNull.Time
	}

	return m, nil
}

// Update implementa la interface producto.Storage
func (p *psqlProduct) Update(m *producto.Modelo) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Nombre,
		stringToNull(m.Detalle),
		m.Precio,
		timeToNull(m.FechaActualizacion),
		m.ID,
	)
	if err != nil {
		return err
	}

	filasAfectadas, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if filasAfectadas == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}

	fmt.Println("se actualizó el producto correctamente")
	return nil
}

// Delete implementa la interface producto.Storage
func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("se eliminó el producto correctamente")
	return nil
}
