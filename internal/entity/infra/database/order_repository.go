package database

import (
	"database/sql"
	"github/gointesivo2/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (or *OrderRepository) Save(order *entity.Order) error {
	_, err := or.Db.Exec("Insert into orders (id, price, tax, final_price) Values(?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil

}

func (or *OrderRepository) GetTotal() (int, error) {
	var total int
	err := or.Db.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
