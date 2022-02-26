package pointers_errors

import "fmt"

type Stringer interface {
	String() string
}

// Esto nos reconoce el tipo Bitcoin como distinto de int
type Bitcoin int

// String() se llama cuando se imprime en métodos de fmt
func (b Bitcoin) String() string {
	// Formatea la string y la devuelve
	return fmt.Sprintf("%d BTC", b)
}

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
