package types

// Money is type to store money amount
type Money int64

// Currency is type to store currencies types
type Currency string

// Currencies codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN is private account number
type PAN string

// Card stores information about cards
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Category represents type of payments
type Category string

// Status represents payments' status
type Status string

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment is a type to represent information about payments
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

// PaymentSource stores type, number and balance of cards
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
