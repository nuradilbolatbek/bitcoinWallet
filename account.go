package bitcoinWallet

type Wallet struct {
	ID      int     `json:"id" db:"id"`
	UserID  int     `json:"user_id" db:"user_id"` // Link to the User
	Balance float64 `json:"balance" db:"balance"` // Balance in bitcoins
}
