package database

import (
	"database/sql"

	"github.com/goFullCycle/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) (int64, error) {
	result, err := r.Db.Exec("Insert into orders (id, price, tax, final_price) values (?,?,?,?)", order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return 0, err
	}

	last, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	return last, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
