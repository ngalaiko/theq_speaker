package types

type Account struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    Gender `json:"gender"`
}
