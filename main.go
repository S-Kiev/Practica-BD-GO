package main

import (
	"log"

	"github.com/S-Kiev/Practica-BD-GO/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProducto := storage.NewPsqlProduct(storage.Pool())
	servicioProducto := producto.NewServicio(storageProducto)

	if err := servicioProducto.Migrate(); err != nil {
		log.Fatalf("migracion del producto: %v", err)
	}

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
