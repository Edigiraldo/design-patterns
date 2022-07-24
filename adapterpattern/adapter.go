package adapterpattern

import "fmt"

type Payment interface {
	Pay()
}

func ProcessPayment(p Payment) {
	fmt.Println("Loading payment...")
	p.Pay()
}

type cashPayment struct {
	amount float64
}

func (p cashPayment) Pay() {
	fmt.Printf("Processing cash payment wiith an amount of %v$USD\n", p.amount)
}

type bankPayment struct {
	amount float64
}

func (b bankPayment) Pay(accountID string) {
	fmt.Printf("Processing bank payment with an amount of %f$USD to %s\n", b.amount, accountID)
}

type bankPaymentAdapter struct {
	bankPayment *bankPayment
	accountID   string
}

func (b bankPaymentAdapter) Pay() {
	b.bankPayment.Pay(b.accountID)
}

func RunExample() {
	cash := cashPayment{amount: 15.56}
	ProcessPayment(cash)

	bank := bankPayment{amount: 935.78}
	bankAdapter := bankPaymentAdapter{bankPayment: &bank, accountID: "testAccountID1234US"}
	ProcessPayment(bankAdapter)
}
