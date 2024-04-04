package repo

import (
	"bitcoinWallet"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) CreateWallet(userID int, wallet bitcoinWallet.Wallet) (int, error) {
	var walletID int
	query := fmt.Sprintf("INSERT INTO %s (user_id, balance) VALUES ($1, $2) RETURNING id", walletsTable)
	err := r.db.QueryRow(query, userID, wallet.Balance).Scan(&walletID)
	if err != nil {
		return 0, err
	}
	return walletID, nil
}

func (r *WalletPostgres) GetWalletByUserID(userID int) (bitcoinWallet.Wallet, error) {
	var wallet bitcoinWallet.Wallet
	query := fmt.Sprintf("SELECT id, user_id, balance FROM %s WHERE user_id = $1", walletsTable)
	err := r.db.Get(&wallet, query, userID)
	return wallet, err
}

func (r *WalletPostgres) DepositToWallet(walletID int, amount float64) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance + $1 WHERE id = $2", walletsTable)
	_, err := r.db.Exec(query, amount, walletID)
	return err
}

func (r *WalletPostgres) WithdrawFromWallet(walletID int, amount float64) error {
	query := fmt.Sprintf("UPDATE %s SET balance = balance - $1 WHERE id = $2", walletsTable)
	_, err := r.db.Exec(query, amount, walletID)
	return err
}

func (r *WalletPostgres) GetWalletBalance(walletID int) (float64, error) {
	var balance float64
	query := fmt.Sprintf("SELECT balance FROM %s WHERE id = $1", walletsTable)
	err := r.db.Get(&balance, query, walletID)
	if err != nil {
		// Handle the error, e.g., if the wallet is not found
		return 0, err
	}

	return balance, nil
}
