package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
	itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
	producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
	"github.com/S-Kiev/Practica-BD-GO/storage"
)

func main() {
	//Iniciar BD
	storage.NewPostgresDB()

	storageProducto := storage.NewPsqlProduct(storage.Pool())
	servicioProducto := producto.NewServicio(storageProducto)

	//Hacer Migracion (Crear Tablas)
	if err := servicioProducto.Migrate(); err != nil {
		log.Fatalf("migracion del producto: %v", err)
	}

	//Insertar Producto (Create)
	p := &producto.Modelo{
		Nombre: "1Kg Arroz",
		Precio: 200,
	}

	if err := servicioProducto.Create(p); err != nil {
		log.Fatalf("insercion de producto: %v", err)
	}

	//Obtener todos los registros (GetAll)

	modelos, err := servicioProducto.GetAll()
	if err != nil {
		log.Fatalf("obtencion de todos los reguistros de producto: %v", err)
	}

	fmt.Println(modelos)

	//Obtener un registro por ID (GetById)

	modelo, err := servicioProducto.GetByID(1)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con ese id")
	case err != nil:
		log.Fatalf("obtencion de del reguistro de producto: %v", err)
	default:
		fmt.Println(modelo)
	}

	//---------------------------------------
	//Para inicializar todas las otras Tablas

	storegeEncabezado := storage.NewPsqlEncabezadoFactura(storage.Pool())
	servicioEncabezado := encabezadofactura.NewServicio(storegeEncabezado)

	if err := servicioEncabezado.Migrate(); err != nil {
		log.Fatalf("migracion del producto: %v", err)
	}

	storageItemFactura := storage.NewPsqlItemFactura(storage.Pool())
	servicioItem := itemfactura.NewServicio(storageItemFactura)

	if err := servicioItem.Migrate(); err != nil {
		log.Fatalf("migracion del item: %v", err)
	}

}
