package types

type Question struct {
	ID      int64   `json:"id"`
	Title   string  `json:"title"`
	Account Account `json:"account"`
}

func (t *Question) String() string {
	return t.Title
}

func (t *Question) Gender() Gender {
	return t.Account.Gender
}
