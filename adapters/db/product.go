package db

import (
	"database/sql"

	"github.com/leandro-koller-bft/hexarch/app"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
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
