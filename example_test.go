package brdoc_test

import (
	"fmt"

	"github.com/danilloteeo/brdoc"
)

func ExampleValidarCNPJ() {
	fmt.Println(brdoc.ValidarCNPJ("11.222.333/0001-81"))
	fmt.Println(brdoc.ValidarCNPJ("12.ABC.345/01DE-35")) // alfanumérico
	fmt.Println(brdoc.ValidarCNPJ("11.222.333/0001-82"))
	// Output:
	// true
	// true
	// false
}

func ExampleFormatarCPF() {
	fmt.Println(brdoc.FormatarCPF("52998224725"))
	// Output: 529.982.247-25
}

func ExampleNormalizarDocumento() {
	fmt.Println(brdoc.NormalizarDocumento("12.abc.345/01de-35"))
	// Output: 12ABC34501DE35
}
