package storage

import (
	"database/sql"
	"fmt"

	producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
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
	mysqlCreateProduct = `INSERT INTO productos(nombre, detalle, precio, fechaCreacion) VALUES(?, ?, ?, ?)`
	mysqlGetAllProduct = `SELECT id, nombre, detalle, precio, 
	fechaCreacion, fechaModificacion FROM productos`
	mysqlGetProductByID = psqlGetAllProduct + " WHERE id = ?"
	mysqlUpdateProduct  = `UPDATE productos SET nombre = ?, detalle = ?,
	precio = ?, fechaModificacion = ? WHERE id = ?`

	/*


		psqlDeleteProduct = `DELETE FROM productos WHERE id = $1`
	*/
)

type scanner interface {
	Scan(dest ...interface{}) error
}

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

// Create implementa la interface Producto.Storage
func (p *MySQLProducto) Create(m *producto.Modelo) error {
	stmt, err := p.db.Prepare(mysqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(
		m.Nombre,
		stringToNull(m.Detalle),
		m.Precio,
		m.FechaCreacion,
	)
	if err != nil {
		return err
	}

	id, err := resultado.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("se creo el producto correctamente con el id: %d\n", m.ID)
	return nil
}

// GetAll implementa la interface product.Storage
func (p *MySQLProducto) GetAll() (producto.Modelos, error) {
	stmt, err := p.db.Prepare(mysqlGetAllProduct)
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

		m, err := scanRowProducto(rows)
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
func (p *MySQLProducto) GetByID(id uint) (*producto.Modelo, error) {
	stmt, err := p.db.Prepare(mysqlGetProductByID)
	if err != nil {
		return &producto.Modelo{}, err
	}
	defer stmt.Close()

	return scanRowProducto(stmt.QueryRow(id))
}

func scanRowProducto(scan scanner) (*producto.Modelo, error) {
	m := &producto.Modelo{}
	var detalle []byte
	fechaActualizacionNull := sql.NullTime{}
	err := scan.Scan(
		&m.ID,
		&m.Nombre,
		&detalle,
		&m.Precio,
		&m.FechaCreacion,
		&fechaActualizacionNull,
	)
	if err != nil {
		return nil, err
	}
	m.Detalle = string(detalle)
	return m, nil
}

// Update implementa la interface producto.Storage
func (p *MySQLProducto) Update(m *producto.Modelo) error {
	stmt, err := p.db.Prepare(mysqlUpdateProduct)
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

/*
VERSIONES ANTERIORES DE SCANNER

func scanRowProductMySQLForRows(row *sql.Rows) (*producto.Modelo, error) {
	m := &producto.Modelo{}
	var detalle []byte
	fechaActualizacionNull := sql.NullTime{}
	err := row.Scan(
		&m.ID,
		&m.Nombre,
		&detalle,
		&m.Precio,
		&m.FechaCreacion,
		&fechaActualizacionNull,
	)
	if err != nil {
		return nil, err
	}
	m.Detalle = string(detalle)
	return m, nil
}

func scanRowProductMySQLForRow(row *sql.Row) (*producto.Modelo, error) {
	m := &producto.Modelo{}
	var detalle []byte
	fechaActualizacionNull := sql.NullTime{}
	err := row.Scan(
		&m.ID,
		&m.Nombre,
		&detalle,
		&m.Precio,
		&m.FechaCreacion,
		&fechaActualizacionNull,
	)
	if err != nil {
		return nil, err
	}
	m.Detalle = string(detalle)
	return m, nil
}

*/
