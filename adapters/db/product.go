package db

import (
	"database/sql"

	"github.com/leandro-koller-bft/hexarch/app"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(id string) (app.IProduct, error) {
	var product app.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product app.IProduct) (app.IProduct, error) {
	var rows int
	err := p.db.QueryRow("SELECT id FROM products WHERE id = ?", product.GetID()).Scan(&rows)

	if err != nil || rows == 0 {
		_, err = p.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDB) create(product app.IProduct) (app.IProduct, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product app.IProduct) (app.IProduct, error) {
	_, err := p.db.Exec(
		"update products set name=?, price=?, status=? where id=?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}
