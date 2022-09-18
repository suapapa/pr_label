package main

type Order struct {
	ID   string `jsong:"id,omitempty"`
	From *Addr  `json:"from"`
	To   *Addr  `json:"to"`
}

type Addr struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PostNumber  string `json:"post_number,omitempty"`
}
