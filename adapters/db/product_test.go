package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/leandro-koller-bft/hexarch/adapters/db"
	"github.com/leandro-koller-bft/hexarch/app"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(sqlDB *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);
	`
	stmt, err := sqlDB.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(sqlDB *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, "disabled")`
	stmt, err := sqlDB.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)
	product, err := productDB.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, app.DISABLED, product.GetStatus())
	require.Equal(t, float64(0), product.GetPrice())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	product := app.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = app.ENABLED

	productResult, err = productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
