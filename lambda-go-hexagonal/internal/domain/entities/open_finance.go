package entities

type Tarifas struct {
	CodigoDenominacao string    `json:"codigoDenominacao"`
	Moeda             string    `json:"moeda"`
	ValorMinimoCobrado float64   `json:"valorMinimoCobrado"`
	ValorMaximoCobrado float64   `json:"valorMaximoCobrado"`
	Quartis           []Quartil `json:"quartis"`
}

type Taxas struct {
	CodigoDenominacao string    `json:"codigoDenominacao"`
	ValorMinimoCobrado float64   `json:"valorMinimoCobrado"`
	ValorMaximoCobrado float64   `json:"valorMaximoCobrado"`
	Quartis           []Quartil `json:"quartis"`
}

type Quartil struct {
	CodigoIdentificador        int     `json:"codigoIdentificador"`
	ValorPercentualMedioCobrado float64 `json:"valorPercentualMedioCobrado"`
	PercentualClientes          int     `json:"percentualClientes"`
}