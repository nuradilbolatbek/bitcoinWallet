package repo

import (
	"bitcoinWallet"
	"github.com/jmoiron/sqlx"
)

type UserAuthentication interface {
	CreateUser(user bitcoinWallet.User) (int, error)
	GetUser(username, password string) (bitcoinWallet.User, error)
}

type WalletRepo interface {
	CreateWallet(userID int, wallet bitcoinWallet.Wallet) (int, error)
	GetWalletByUserID(userID int) (bitcoinWallet.Wallet, error)
	DepositToWallet(walletID int, amount float64) error
	WithdrawFromWallet(walletID int, amount float64) error
}

type Repository struct {
	UserAuthentication
	WalletRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserAuthentication: NewAuthPostgres(db),
		WalletRepo:         NewWalletPostgres(db),
	}
}
