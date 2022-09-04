package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//https://en.wikipedia.org/wiki/Golomb_coding
//http://multimedia.ufp.pt/codecs/compressao-sem-perdas/codificacao-estatistica/codificacao-de-golomb-rice/

//Passo a Passo Golomb-rice

// 1- Converter o dígito para decimal (isso já é feito na leitura do arquivo)
// 2 - Dividir o valor decimal pelo divisor selecionado gerando Y = Resto da Divisão, X = Resultado da Divisão
// 3 - Gerar o valor codificado X zeros, stop bit(1), Y representado em binário pela quantidade de bits resultante de logaritmo(divisor)

func GolombCoding(divisor int, dados []byte) {

	if (divisor & (divisor - 1)) != 0 {
		panic("o divisor nao é uma potencia de 2")
	}

	for _, x := range dados {
		//2
		valorInt := int(x)

		resultadoDivisao := valorInt / divisor
		restoDivisao := valorInt % divisor

		fmt.Println(resultadoDivisao, restoDivisao)
	}

}

func main() {

	data, err := os.ReadFile("files/test.txt")
	check(err)

	GolombCoding(4, data)

}
