package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID        int
	IdAdmin   int
	Name      string
	Unit      string
	Stock     int
	Price     int
	Image     string
	CreatedAt time.Time
}

//logic
type ProductUseCase interface {
	CreateProduct(newProduct Product, idAdmin int) int
	GetAllProduct(limit, offset int) (data []Product, err error)
	UpdateProduct(updatedData Product, idProduct int) int
	DeleteProduct(productid int) int
}

//query
type ProductData interface {
	CreateProductData(newProduct Product) Product
	GetAllProductData(limit, offset int) (data []Product, err error)
	UpdateProductData(updatedData Product) Product
	DeleteProductData(productid int) (row int, err error)
}

//handler
type ProductHandler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}
