package main

import (
	"log"

	encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
	itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
	producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
	"github.com/S-Kiev/Practica-BD-GO/storage"
)

func main() {
	//Iniciar BD
	storage.NewPostgresDB()

	//Hacer Migracion (Crear Tablas)
	storageProducto := storage.NewPsqlProduct(storage.Pool())
	servicioProducto := producto.NewServicio(storageProducto)

	if err := servicioProducto.Migrate(); err != nil {
		log.Fatalf("migracion del producto: %v", err)
	}

	//Insertar Producto (Create)
	m := &producto.Modelo{
		Nombre: "1Kg Arroz",
		Precio: 200,
	}

	if err := servicioProducto.Create(m); err != nil {
		log.Fatalf("migracion del producto: %v", err)
	}

	//---------------------------------------
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
