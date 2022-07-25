package app

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type IProductService interface {
	Get(id string) (IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}

type IProductReader interface {
	Get(id string) (IProduct, error)
}

type IProductWriter interface {
	Save(product IProduct) (IProduct, error)
}

type IProductPersistence interface {
	IProductReader
	IProductWriter
}

const (
	DISABLED              = "disabled"
	ENABLED               = "enabled"
	INVALID_PRICE_ERROR   = "the price must be greater than zero to enable the product"
	NEGATIVE_PRICE_ERROR  = "the price must be greater or equal than zero"
	INVALID_DISABLE_ERROR = "the price must be zero to disable the product"
	INVALID_STATUS_ERROR  = "the status must be enabled or disabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"optional"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {

	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New(INVALID_STATUS_ERROR)
	}

	if p.Price < 0 {
		return false, errors.New(NEGATIVE_PRICE_ERROR)
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New(INVALID_PRICE_ERROR)
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New(INVALID_DISABLE_ERROR)
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
