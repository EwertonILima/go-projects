package entities

type KafkaEntity struct {
    Data Data `json:"data"`
}

type Data struct {
    SpecVersion         string                `json:"spec_version"`
    Type                string                `json:"type"`
    Source              string                `json:"source"`
    Id                  string                `json:"id"`
    Time                string                `json:"time"`
    EventVersion        string                `json:"event_version"`
    TransactionID       string                `json:"transaction_id"`
    DataContentType     string                `json:"data_content_type"`
    EventoKafkaProdutoDto EventoKafkaProdutoDto `json:"evento_kafka_produto_dto"`
}

type EventoKafkaProdutoDto struct {
    Dn                      int                 `json:"dn"`
    CodigoEmpresa           string              `json:"codigo_empresa"`
    NomeProduto             string              `json:"nome_produto"`
    Variante                string                 `json:"variante"`
    Bandeira                string                  `json:"bandeira"`
    TipoPessoa              string              `json:"tipo_pessoa"`
    UriProgramaRecompensa   string              `json:"uri_programa_recompensa"`
    CriteriosElegibilidade  string              `json:"criterios_elegibilidade"`
    RegrasPagamentoMinimo   []RegraPagamentoMinimo `json:"regras_pagamento_minimo"`
    Tarifas                 []Tarifas           `json:"tarifas"`
    Taxas                   []Taxas             `json:"taxas"`
}

type RegraPagamentoMinimo struct {
    CodigoRegra string `json:"codigo_regra"`
}