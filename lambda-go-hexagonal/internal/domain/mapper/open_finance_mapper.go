package mapper

import (
	"time"
	"lambda-go-hexagonal/internal/domain/entities"

	"github.com/google/uuid"
)

func MapSqsToOpenFinanceJson(sqsMessage entities.SqsEntity) entities.KafkaEntity {
    kafkaMessage := entities.KafkaEntity{
        Data: entities.Data{
            SpecVersion: "1.0",
            Type: "open-banking",
            Source: "configuration-manager",
            Id: uuid.New().String(),
            Time: time.Now().Format("15:04:05"),
            EventVersion: "1.0",
            TransactionID: uuid.New().String(),
            DataContentType: "application/json",
            EventoKafkaProdutoDto: entities.EventoKafkaProdutoDto{
                Dn: sqsMessage.Dn,
                CodigoEmpresa: "",
                NomeProduto: "",
                Variante: "",
                Bandeira: "",
                TipoPessoa: "",
                UriProgramaRecompensa: "",
                CriteriosElegibilidade: "",
                RegrasPagamentoMinimo: []entities.RegraPagamentoMinimo{
                    {CodigoRegra: "PAGMINDESLIGADO"},
                },
                Tarifas: sqsMessage.Tarifas,
                Taxas: sqsMessage.Taxas,
            },
        },
    }

    return kafkaMessage
}