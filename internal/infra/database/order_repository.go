package database

import (
	"database/sql"
	"fmt"

	"github.com/EuricoCruz/cleanarch_challenge/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) List(limit int, offset int) ([]*entity.Order, error) {
	fmt.Printf("Listando pedidos com limit %d e offset %d\n", limit, offset)
	query := "Select * from orders LIMIT ? OFFSET ?"
	rows, err := r.Db.Query(query, limit, offset); if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer fmt.Println("conexão fechada")

	var orders []*entity.Order

	for rows.Next() {
		order := &entity.Order{}
		err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	fmt.Println("Retornando os dados")

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}