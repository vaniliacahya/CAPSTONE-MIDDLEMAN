package data

import (
	"middleman-capstone/domain"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Qty       int
	Status    string
	Subtotal  int
	UserID    int
	ProductID int
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Role     string
	Address  string
	Product  []Product
}

type Product struct {
	gorm.Model
	Stock  int
	Status string
	Name   string
	Image  string
	Unit   string
	Price  int
	UserID int
	Cart   []Cart `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   User
}

func (c *Cart) ToDomain() domain.Cart {
	return domain.Cart{
		ID:        int(c.ID),
		Qty:       c.Qty,
		Status:    c.Status,
		UserID:    c.UserID,
		Subtotal:  c.Subtotal,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Product: domain.ProductCart{
			ID:           int(c.Product.ID),
			ProductName:  c.Product.Name,
			Unit:         c.Product.Unit,
			Stock:        c.Product.Stock,
			Price:        c.Product.Price,
			ProductImage: c.Product.Image,
		},
	}

}

func ParseToArr(arr []Cart) []domain.Cart {
	var res []domain.Cart
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}
func FromDomain(data domain.Cart) Cart {
	var res Cart
	res.Qty = data.Qty
	res.Status = data.Status
	res.UserID = data.UserID
	res.ProductID = data.Product.ID
	res.Subtotal = data.Subtotal
	return res
}