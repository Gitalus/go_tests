package pointers_errors

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// No es necesario este puntero pero por buena práctica para mantener consistencia
func (w *Wallet) Balance() int {
	return w.balance
}
