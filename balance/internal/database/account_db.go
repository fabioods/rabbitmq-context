package database

import (
	"database/sql"
	"fmt"
	"github.com/fabioods/balance/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	fmt.Println("db ", a.DB)
	stmt, err := a.DB.Prepare("select a.id, a.client_id, a.balance, a.created_at FROM accounts a WHERE a.id = ?")
	defer stmt.Close()
	if err != nil {
		fmt.Println("Error to prepare statement")
		return nil, err
	}

	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID,
		&account.ClientID,
		&account.Balance,
		&account.CreatedAt)
	if err != nil {
		fmt.Println("Error to execute query")
		return nil, err
	}
	return &account, nil
}

func (a *AccountDB) UpdateBalance(account *entity.Account) error {
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}
	return nil
}
