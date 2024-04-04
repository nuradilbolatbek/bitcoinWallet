package service

import (
	"bitcoinWallet"
	"bitcoinWallet/package/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserAuthentication interface {
	CreateUser(user bitcoinWallet.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type WalletManagement interface {
	CreateWallet(userID int, wallet bitcoinWallet.Wallet) (int, error)
	GetWalletByUserID(userID int) (bitcoinWallet.Wallet, error)
	DepositToWallet(walletID int, amount float64) error
	WithdrawFromWallet(walletID int, amount float64) error
}

type Service struct {
	UserAuthentication
	WalletManagement
}

func NewService(repos *repo.Repository) *Service {
	return &Service{

		UserAuthentication: NewAuthService(repos.UserAuthentication), // Assuming you have a user repository
		WalletManagement:   NewWalletService(repos.WalletRepo),       // Assuming you have a wallet repository
	}
}
