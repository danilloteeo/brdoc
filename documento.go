package brdoc

import (
	"strings"
	"unicode"
)

// NormalizarDocumento prepara CPF e CNPJ para validar e gravar: tira
// pontuação e espaços e põe as letras em maiúsculas. Letras ficam porque o
// CNPJ alfanumérico (IN RFB 2.229/2024), em vigor desde julho de 2026, as
// tem nas doze primeiras posições — "12.abc.345/01de-35" vira "12ABC34501DE35".
func NormalizarDocumento(valor string) string {
	var b strings.Builder
	b.Grow(len(valor))
	for _, r := range valor {
		switch {
		case r >= '0' && r <= '9', r >= 'A' && r <= 'Z':
			b.WriteRune(r)
		case r >= 'a' && r <= 'z':
			b.WriteRune(r - 'a' + 'A')
		}
	}
	return b.String()
}

func SomenteDigitos(valor string) string {
	var b strings.Builder
	b.Grow(len(valor))
	for _, r := range valor {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func ValidarCPF(valor string) bool {
	cpf := NormalizarDocumento(valor)
	if len(cpf) != 11 || !apenasDigitos(cpf) || todosIguais(cpf) {
		return false
	}
	return digitoVerificador(cpf[:9], 10) == rune(cpf[9]) &&
		digitoVerificador(cpf[:10], 11) == rune(cpf[10])
}

// ValidarCNPJ aceita o CNPJ numérico e o alfanumérico: doze posições de
// dígitos ou letras maiúsculas e dois verificadores numéricos. O cálculo é o
// módulo 11 de sempre, com cada caractere valendo o código ASCII menos 48
// ('0' = 0, 'A' = 17) — para o CNPJ só de números, nada muda.
func ValidarCNPJ(valor string) bool {
	cnpj := NormalizarDocumento(valor)
	if len(cnpj) != 14 || !apenasDigitos(cnpj[12:]) || todosIguais(cnpj) {
		return false
	}
	pesosPrimeiro := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesosSegundo := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	return digitoComPesos(cnpj[:12], pesosPrimeiro) == rune(cnpj[12]) &&
		digitoComPesos(cnpj[:13], pesosSegundo) == rune(cnpj[13])
}

// ChaveNoFormato confere o desenho da chave de acesso de NF-e/NFC-e, sem o
// verificador: 44 posições, todas dígitos menos as doze primeiras do CNPJ do
// emitente (posições 7 a 18), que podem ter letras desde o CNPJ alfanumérico
// (NT 2025.001).
func ChaveNoFormato(chave string) bool {
	if len(chave) != 44 {
		return false
	}
	for i := 0; i < len(chave); i++ {
		c := chave[i]
		letraNoCNPJ := i >= 6 && i < 18 && c >= 'A' && c <= 'Z'
		if !letraNoCNPJ && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}

func ValidarCPFouCNPJ(valor string) bool {
	switch len(NormalizarDocumento(valor)) {
	case 11:
		return ValidarCPF(valor)
	case 14:
		return ValidarCNPJ(valor)
	default:
		return false
	}
}

func ValidarCEP(valor string) bool {
	return len(SomenteDigitos(valor)) == 8
}

func ValidarCodigoIBGE(valor string) bool {
	return len(SomenteDigitos(valor)) == 7
}

func ValidarNCM(valor string) bool {
	return len(SomenteDigitos(valor)) == 8
}

func ValidarGTIN(valor string) bool {
	gtin := SomenteDigitos(valor)
	switch len(gtin) {
	case 8, 12, 13, 14:
	default:
		return false
	}
	if todosIguais(gtin) {
		return false
	}

	soma := 0
	peso := 3
	for i := len(gtin) - 2; i >= 0; i-- {
		soma += int(gtin[i]-'0') * peso
		if peso == 3 {
			peso = 1
		} else {
			peso = 3
		}
	}
	verificador := (10 - soma%10) % 10
	return verificador == int(gtin[len(gtin)-1]-'0')
}

func FormatarCPF(valor string) string {
	cpf := NormalizarDocumento(valor)
	if len(cpf) != 11 {
		return valor
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

func FormatarCNPJ(valor string) string {
	cnpj := NormalizarDocumento(valor)
	if len(cnpj) != 14 {
		return valor
	}
	return cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]
}

func FormatarCEP(valor string) string {
	cep := SomenteDigitos(valor)
	if len(cep) != 8 {
		return valor
	}
	return cep[:5] + "-" + cep[5:]
}

func apenasDigitos(valor string) bool {
	for i := 0; i < len(valor); i++ {
		if valor[i] < '0' || valor[i] > '9' {
			return false
		}
	}
	return true
}

func todosIguais(valor string) bool {
	for i := 1; i < len(valor); i++ {
		if valor[i] != valor[0] {
			return false
		}
	}
	return true
}

func digitoVerificador(base string, pesoInicial int) rune {
	soma := 0
	for i, r := range base {
		soma += int(r-'0') * (pesoInicial - i)
	}
	return digitoDoResto(soma)
}

func digitoComPesos(base string, pesos []int) rune {
	soma := 0
	for i, r := range base {
		soma += int(r-'0') * pesos[i]
	}
	return digitoDoResto(soma)
}

func digitoDoResto(soma int) rune {
	resto := soma % 11
	if resto < 2 {
		return '0'
	}
	return rune('0' + (11 - resto))
}
