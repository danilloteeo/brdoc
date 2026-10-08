package brdoc

import "testing"

func TestValidarCPF(t *testing.T) {
	validos := []string{"529.982.247-25", "52998224725", "111.444.777-35"}
	for _, v := range validos {
		if !ValidarCPF(v) {
			t.Errorf("ValidarCPF(%q) = false, esperado true", v)
		}
	}

	invalidos := []string{"", "123", "111.111.111-11", "529.982.247-26", "5299822472a", "00000000000"}
	for _, v := range invalidos {
		if ValidarCPF(v) {
			t.Errorf("ValidarCPF(%q) = true, esperado false", v)
		}
	}
}

func TestValidarCNPJ(t *testing.T) {
	validos := []string{"11.222.333/0001-81", "11222333000181"}
	for _, v := range validos {
		if !ValidarCNPJ(v) {
			t.Errorf("ValidarCNPJ(%q) = false, esperado true", v)
		}
	}

	invalidos := []string{"", "11222333000182", "11.111.111/1111-11", "1122233300018"}
	for _, v := range invalidos {
		if ValidarCNPJ(v) {
			t.Errorf("ValidarCNPJ(%q) = true, esperado false", v)
		}
	}
}

func TestValidarCPFouCNPJ(t *testing.T) {
	if !ValidarCPFouCNPJ("52998224725") {
		t.Error("CPF válido rejeitado")
	}
	if !ValidarCPFouCNPJ("11222333000181") {
		t.Error("CNPJ válido rejeitado")
	}
	if ValidarCPFouCNPJ("123456") {
		t.Error("documento com tamanho inválido aceito")
	}
}

func TestValidarGTIN(t *testing.T) {
	validos := []string{"7891000315507", "7894900011517", "78915817"}
	for _, v := range validos {
		if !ValidarGTIN(v) {
			t.Errorf("ValidarGTIN(%q) = false, esperado true", v)
		}
	}

	invalidos := []string{"", "7891000315508", "12345", "0000000000000"}
	for _, v := range invalidos {
		if ValidarGTIN(v) {
			t.Errorf("ValidarGTIN(%q) = true, esperado false", v)
		}
	}
}

func TestSomenteDigitos(t *testing.T) {
	if got := SomenteDigitos("529.982.247-25"); got != "52998224725" {
		t.Errorf("SomenteDigitos() = %q", got)
	}
}

func TestFormatadores(t *testing.T) {
	if got := FormatarCPF("52998224725"); got != "529.982.247-25" {
		t.Errorf("FormatarCPF() = %q", got)
	}
	if got := FormatarCNPJ("11222333000181"); got != "11.222.333/0001-81" {
		t.Errorf("FormatarCNPJ() = %q", got)
	}
	if got := FormatarCEP("01001000"); got != "01001-000" {
		t.Errorf("FormatarCEP() = %q", got)
	}
	if got := FormatarCPF("123"); got != "123" {
		t.Errorf("entrada inválida deve ser devolvida intacta, veio %q", got)
	}
}

// O CNPJ alfanumérico (IN RFB 2.229/2024) vale desde julho de 2026. O exemplo
// é o da própria Receita.
func TestValidarCNPJAlfanumerico(t *testing.T) {
	for _, v := range []string{"12.ABC.345/01DE-35", "12ABC34501DE35", "12.abc.345/01de-35"} {
		if !ValidarCNPJ(v) {
			t.Errorf("ValidarCNPJ(%q) = false, esperado true", v)
		}
	}
	// Verificador errado, letra nos verificadores, caractere fora de [0-9A-Z].
	for _, v := range []string{"12ABC34501DE36", "12ABC34501DE3A", "12ABC34501DÉ35", "AAAAAAAAAAAAAA"} {
		if ValidarCNPJ(v) {
			t.Errorf("ValidarCNPJ(%q) = true, esperado false", v)
		}
	}
	if !ValidarCPFouCNPJ("12.ABC.345/01DE-35") {
		t.Error("CNPJ alfanumérico rejeitado em ValidarCPFouCNPJ")
	}
	if ValidarCPFouCNPJ("5299822472A") {
		t.Error("CPF não tem letra")
	}
}

func TestNormalizarDocumento(t *testing.T) {
	casos := map[string]string{
		"12.abc.345/01de-35":   "12ABC34501DE35",
		"529.982.247-25":       "52998224725",
		" 11.222.333/0001-81 ": "11222333000181",
		"":                     "",
	}
	for entrada, esperado := range casos {
		if obtido := NormalizarDocumento(entrada); obtido != esperado {
			t.Errorf("NormalizarDocumento(%q) = %q, esperado %q", entrada, obtido, esperado)
		}
	}
	if obtido := FormatarCNPJ("12abc34501de35"); obtido != "12.ABC.345/01DE-35" {
		t.Errorf("FormatarCNPJ alfanumérico = %q", obtido)
	}
}
