package restaurant

import (
	"database/sql"
	"fmt"
)

type Menu struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Order struct {
	Id          int `json:"id"`
	TableNamber int `json:"tableNamber"`
}

type Item struct {
	ItemId   int `json:"id"`
	Quantity int `json:"tableNamber"`
}

type Detail struct {
	OrderId    int     `json:"orderId"`
	TotalPrice float64 `json:"totalPrice"`
	WithoutTax float64 `json:"withoutTax"`
	Tips       float64 `json:"tips"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListMenu() ([]*Menu, error) {
	rows, err := s.Db.Query("SELECT * FROM menu")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menu []*Menu
	for rows.Next() {
		m := Menu{}
		err := rows.Scan(&m.Id, &m.Name, &m.Price)
		if err != nil {
			return nil, err
		}
		menu = append(menu, &m)
	}

	if menu == nil {
		menu = make([]*Menu, 0)
	}
	return menu, nil
}

func (s *Store) CreateOrder(table_number int, items []*Item) (*Detail, error) {
	if table_number <= 0 {
		return nil, fmt.Errorf("Table number should be positive.")
	}

	orderRow := s.Db.QueryRow("INSERT INTO orders (\"table_number\") VALUES ($1) RETURNING id", table_number)

	var orderId int
	err := orderRow.Scan(&orderId)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		res := s.Db.QueryRow("SELECT id FROM menu WHERE id = $1", item.ItemId)
		var itemId int
		if err := res.Scan(&itemId); err != nil {
			return nil, fmt.Errorf("Item not supported: %d", item.ItemId)
		}
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("Quantity of items should more than 0: %d", item.ItemId)
		}

		s.Db.QueryRow("INSERT INTO order_details (order_id, menu_id, quantity) VALUES ($1, $2, $3)", orderId, item.ItemId, item.Quantity)
	}

	var totalPrice float64 = 0.0
	var withoutTax float64 = 0.0
	var tips float64 = 0.0

	for _, item := range items {
		res := s.Db.QueryRow("SELECT id FROM menu WHERE id = $1", item.ItemId)

		var Name string
		var Price int

		if err := res.Scan(&Price, &Name); err != nil {
			return nil, fmt.Errorf("Item not supported: %d", item.ItemId)
		}

		withoutTax = withoutTax + float64((Price * item.Quantity))
		totalPrice = withoutTax * (1 - 0.05)
		tips = totalPrice * 0.01
	}

	detail := &Detail{
		OrderId:    orderId,
		TotalPrice: totalPrice,
		WithoutTax: withoutTax,
		Tips:       tips,
	}

	return detail, nil
}
