package main

import (
	producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
	"github.com/S-Kiev/Practica-BD-GO/storage"
)

/*"database/sql"
"errors"
"fmt"

encabezadofactura "github.com/S-Kiev/Practica-BD-GO/pkg/EncabezadoFactura"
factura "github.com/S-Kiev/Practica-BD-GO/pkg/Factura"
itemfactura "github.com/S-Kiev/Practica-BD-GO/pkg/ItemFactura"
producto "github.com/S-Kiev/Practica-BD-GO/pkg/Producto"
"github.com/S-Kiev/Practica-BD-GO/storage"
*/

func main() {
	storage.NewMySQLDB()

	storageProducto := storage.NewMySQLProducto(storage.Pool())
	servicioProducto := producto.NewServicio(storageProducto)

	/*
			//Hacer Migracion (Crear Tablas)
			if err := servicioProducto.Migrate(); err != nil {
				log.Fatalf("migracion del producto: %v", err)
			}

			storegeEncabezado := storage.NewMySQLEncabezadoFactura(storage.Pool())
			servicioEncabezado := encabezadofactura.NewServicio(storegeEncabezado)

			if err := servicioEncabezado.Migrate(); err != nil {
				log.Fatalf("migracion del producto: %v", err)
			}

			storageItemFactura := storage.NewMySQLItemFactura(storage.Pool())
			servicioItem := itemfactura.NewServicio(storageItemFactura)

			if err := servicioItem.Migrate(); err != nil {
				log.Fatalf("migracion del item: %v", err)
			}


		//Insertar Producto (Create)
		p := &producto.Modelo{
			Nombre: "Papas Fritas",
			Precio: 400,
		}

		if err := servicioProducto.Create(p); err != nil {
			log.Fatalf("insercion de producto: %v", err)
		}
	*/

}

/*
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

	//Para actualizar registros (Update)

	//A partir de aqui como err ya fue declarada ya no se usa:=
	//sino que para los ejemplos subsiguientes se asigna =
	modelo = &producto.Modelo{
		ID:      2,
		Nombre:  "1Kg Papa",
		Detalle: "Alta en almidon",
		Precio:  20,
	}

	err = servicioProducto.Update(modelo)
	if err != nil {
		log.Fatalf("actualizacion del registro de producto: %v", err)
	}

	//Para eliminar reguistro (Delete)

	err = servicioProducto.Delete(3)
	if err != nil {
		log.Fatalf("eliminación del registro de producto: %v", err)
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

	storageFactura := storage.NewPsqlFactura(
		storage.Pool(),
		storegeEncabezado,
		storageItemFactura,
	)

	//El Modelo de factura tiene un encabezado (el cual a su vez en este caso solo requiere del nombre del cliente)
	//Y los items que contiene, que es un slice de item
	facturaPrueba := &factura.Modelo{
		Encabezado: &encabezadofactura.Modelo{
			Cliente: "Ezequiel Viera",
		},
		Items: itemfactura.Modelos{
			&itemfactura.Modelo{ProductoID: 1},
			&itemfactura.Modelo{ProductoID: 2},
		},
	}

	servicioFactura := factura.NewService(storageFactura)
	if err := servicioFactura.Create(facturaPrueba); err != nil {
		log.Fatalf("error al crear la factura: %v", err)
	}
*/
