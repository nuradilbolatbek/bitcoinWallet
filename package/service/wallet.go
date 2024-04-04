package service

import (
	"bitcoinWallet"
	"bitcoinWallet/package/repo"
	"sync"
)

type WalletService struct {
	repo repo.WalletRepo
	mu   sync.Mutex
}

func NewWalletService(repo repo.WalletRepo) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(userID int, wallet bitcoinWallet.Wallet) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.CreateWallet(userID, wallet)
}

func (s *WalletService) GetWalletByUserID(userID int) (bitcoinWallet.Wallet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.GetWalletByUserID(userID)
}

func (s *WalletService) DepositToWallet(walletId int, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.DepositToWallet(walletId, amount)
}

func (s *WalletService) WithdrawFromWallet(walletId int, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.WithdrawFromWallet(walletId, amount)
}
