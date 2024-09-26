package data

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Account struct {
	Id           int64           `json:"id"`
	FirstName    string          `json:"first_name"`
	LastName     string          `json:"last_name"`
	Email        string          `json:"email"`
	Birthday     time.Time       `json:"birthday"`
	CreationDate time.Time       `json:"creation_date"`
	Balance      decimal.Decimal `json:"balance"`
	Version      int32           `json:"version"`
}

type AccountModel struct {
	DB *sql.DB
}

func (m AccountModel) Insert(account *Account) error {
	query := `
insert into accounts (first_name, last_name, email, birthday)
values ($1, $2, $3, $4)
returning id, creation_date, balance, version;`

	args := []any{account.FirstName, account.LastName, account.Email, account.Birthday}
	return m.DB.QueryRow(query, args...).Scan(&account.Id, &account.CreationDate, &account.Balance, &account.Version)
}
