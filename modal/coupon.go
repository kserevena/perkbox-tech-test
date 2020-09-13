package modal

import (
	"database/sql"
	"errors"
	"fmt"
)

type Coupon struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	Value     int    `json:"value"`
	CreatedAt string `json:"createdAt"`
	Expiry    string `json:"expiry"`
}

func (c *Coupon) GetCoupons(db *sql.DB) ([]Coupon, error) {
	stmt := "SELECT * FROM COUPON"
	rows, err := db.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	coupons := []Coupon{}

	for rows.Next() {
		var c Coupon
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Brand,
			&c.Value,
			&c.CreatedAt,
			&c.Expiry,
		); err != nil {
			return nil, err
		}
		coupons = append(coupons, c)
	}

	return coupons, nil
}

func (c *Coupon) GetCoupon(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (c *Coupon) CreateCoupon(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO coupon(name, brand, value, createdAt, expiry) VALUES('%s', '%s', %d, '%s', '%s')",
		c.Name,
		c.Brand,
		c.Value,
		c.CreatedAt,
		c.Expiry,
	)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}
