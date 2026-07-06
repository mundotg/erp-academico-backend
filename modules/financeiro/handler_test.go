package financeiro

import (
	"testing"

	"erp-academico-backend/models"
)

func TestMontarValidacaoDivida(t *testing.T) {
	resultado := montarValidacaoDivida(
		7,
		[]models.Pagamento{{Valor: 1000}, {Valor: 2500}},
		[]models.Encargo{{Valor: 500}},
	)

	if resultado.AlunoID != 7 {
		t.Fatalf("AlunoID = %d, esperado 7", resultado.AlunoID)
	}
	if !resultado.PossuiDivida {
		t.Fatal("esperava possuir dívida")
	}
	if resultado.TotalPagamentos != 3500 {
		t.Fatalf("TotalPagamentos = %.2f, esperado 3500", resultado.TotalPagamentos)
	}
	if resultado.TotalEncargos != 500 {
		t.Fatalf("TotalEncargos = %.2f, esperado 500", resultado.TotalEncargos)
	}
	if resultado.TotalEmDivida != 4000 {
		t.Fatalf("TotalEmDivida = %.2f, esperado 4000", resultado.TotalEmDivida)
	}
}
