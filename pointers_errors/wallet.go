package pointers_errors

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// No es necesario este puntero pero por buena práctica para mantener consistencia
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
