package theq_speak

type Config struct {
	ApiKey string `yaml:"ApiKey"`
}

type Question struct {
	Id      int64   `json:"id"`
	Title   string  `json:"title"`
	Account Account `json:"account"`
}

type Account struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
}
