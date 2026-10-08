# brdoc

[![test](https://github.com/danilloteeo/brdoc/actions/workflows/test.yml/badge.svg)](https://github.com/danilloteeo/brdoc/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/danilloteeo/brdoc.svg)](https://pkg.go.dev/github.com/danilloteeo/brdoc)

Validação e formatação de documentos brasileiros em Go, **sem dependências**.

Suporta o **CNPJ alfanumérico** (IN RFB 2.229/2024), em vigor desde julho de 2026.

Extraído de um PDV em produção que emite NF-e e NFC-e.

## Instalação

```sh
go get github.com/danilloteeo/brdoc
```

## Uso

```go
import "github.com/danilloteeo/brdoc"

brdoc.ValidarCPF("529.982.247-25")        // true
brdoc.ValidarCNPJ("11.222.333/0001-81")   // true
brdoc.ValidarCNPJ("12.ABC.345/01DE-35")   // true (alfanumérico)
brdoc.ValidarCPFouCNPJ("52998224725")     // true

brdoc.FormatarCNPJ("12abc34501de35")      // "12.ABC.345/01DE-35"
brdoc.NormalizarDocumento("529.982.247-25") // "52998224725"

brdoc.ValidarGTIN("7891000315507")        // true (EAN-13)
```

## O que tem

| Função | O que faz |
|---|---|
| `ValidarCPF` | Tamanho, dígitos repetidos e verificadores |
| `ValidarCNPJ` | Numérico e alfanumérico, módulo 11 |
| `ValidarCPFouCNPJ` | Decide pelo tamanho |
| `ValidarGTIN` | GTIN-8, 12, 13 e 14 (EAN/UPC) com verificador |
| `ValidarCEP`, `ValidarNCM`, `ValidarCodigoIBGE` | Quantidade de dígitos |
| `ChaveNoFormato` | Formato da chave de acesso de NF-e/NFC-e (44 posições, letras no CNPJ) |
| `FormatarCPF`, `FormatarCNPJ`, `FormatarCEP` | Máscaras; entrada inválida volta intacta |
| `NormalizarDocumento`, `SomenteDigitos` | Limpeza para validar e gravar |

## CNPJ alfanumérico

As 12 primeiras posições podem ter letras `A-Z`; os 2 verificadores continuam numéricos. O cálculo é o mesmo módulo 11, com cada caractere valendo seu código ASCII menos 48 (`'0'` = 0, `'A'` = 17). Para CNPJs só com números, o resultado é idêntico ao de antes.

## Licença

MIT
