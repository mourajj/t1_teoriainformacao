package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

func GolombCoding(divisor int, dados []byte) []string {

	result := []string{}

	if (divisor & (divisor - 1)) != 0 {
		panic("o divisor nao é uma potencia de 2")
	}

	for _, x := range dados {
		//2
		output := ""
		valorInt := int(x)
		resultadoDivisao := valorInt / divisor
		restoDivisao := valorInt % divisor

		fmt.Printf("Resultado divisao: %d, Resto divisao: %d\n", resultadoDivisao, restoDivisao)

		log := math.Log2(float64(divisor))

		for i := 0; i < resultadoDivisao; i++ {
			output = output + "0"
		}

		var sufix string = strconv.FormatInt(int64(restoDivisao), 2)
		fmt.Printf("%d representado em binário: %s\n", restoDivisao, sufix)
		fmt.Printf("esse numero precisa ser representado em %d bits após o stop bit (1)\n", int(log))

		for float64(len(sufix)) < log {
			sufix = "0" + sufix
		}

		output = output + "1" + sufix
		fmt.Println(sufix)
		result = append(result, output)
	}
	return result

}

func main() {

	data, err := os.ReadFile("files/test.txt")
	check(err)

	fmt.Println(data)

	fmt.Println(GolombCoding(4, data))

}
