package encoding

import (
	"fmt"
	"math"
	"strconv"
)

func Golomb(divisor int, dados []byte) []string {

	result := []string{}

	if (divisor & (divisor - 1)) != 0 {
		panic("o divisor nao é uma potencia de 2")
	}

	for _, x := range dados {
		output := ""
		valorInt := int(x)
		resultadoDivisao := valorInt / divisor
		restoDivisao := valorInt % divisor

		//fmt.Printf("Resultado divisao: %d, Resto divisao: %d\n", resultadoDivisao, restoDivisao)
		log := math.Log2(float64(divisor))

		for i := 0; i < resultadoDivisao; i++ {
			output = output + "0"
		}

		var sufix string = strconv.FormatInt(int64(restoDivisao), 2)
		//fmt.Printf("%d representado em binário: %s\n", restoDivisao, sufix)
		//fmt.Printf("esse numero precisa ser representado em %d bits após o stop bit (1)\n", int(log))

		for float64(len(sufix)) < log {
			sufix = "0" + sufix
		}

		output = output + "1" + sufix
		//fmt.Println(sufix)
		result = append(result, output)
	}
	return result

}

func Unario(dados []byte) []string {
	result := []string{}

	for _, x := range dados {

		output := ""
		valorInt := int(x)

		for i := 0; i < valorInt; i++ {
			output = output + "0"
		}
		output = output + "1"
		result = append(result, output)

	}

	return result
}

func EliasGamma(dados []byte) []string {

	result := []string{}

	for _, x := range dados {
		log := int(math.Log2(float64(x)))
		resto := int(float64(x) - math.Pow(2, float64(log)))
		output := ""

		fmt.Printf("prefixo: %d, sufixo: %d\n", log, resto)

		for i := 0; i < log; i++ {
			output = output + "0"
		}

		var sufix string = strconv.FormatInt(int64(resto), 2)

		for int(len(sufix)) < log {
			sufix = "0" + sufix
		}
		output = output + "1" + sufix
		result = append(result, output)
	}

	return result
}
