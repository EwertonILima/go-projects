package entities

type SqsEntity struct {
	Dn      int       `json:"dn"`
	Tarifas []Tarifas `json:"tarifas"`
	Taxas   []Taxas   `json:"taxas"`
}