package encoding

import (
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
	contem0 := false

	//Verificando de tem um 0 nos dados, caso tenha, adiciona +1 em cada unidade do slice de dados.
	for _, y := range dados {
		if y == 0 {
			contem0 = true
			break
		}
	}

	if contem0 {
		for _, x := range dados {
			x = x + 1
			log := int(math.Log2(float64(x)))
			resto := int(float64(x) - math.Pow(2, float64(log)))
			output := ""

			for i := 0; i < log; i++ {
				output = output + "0"
			}

			var sufix string = strconv.FormatInt(int64(resto), 2)

			for int(len(sufix)) < log {
				sufix = "0" + sufix
			}
			output = output + "1" + sufix

			if x == 1 {
				output = "1"
			}
			result = append(result, output)

		}
	} else {
		for _, x := range dados {
			log := int(math.Log2(float64(x)))
			resto := int(float64(x) - math.Pow(2, float64(log)))
			output := ""

			for i := 0; i < log; i++ {
				output = output + "0"
			}

			var sufix string = strconv.FormatInt(int64(resto), 2)

			for int(len(sufix)) < log {
				sufix = "0" + sufix
			}
			output = output + "1" + sufix

			if x == 1 {
				output = "1"
			}
			result = append(result, output)
		}
	}

	return result
}

func closestFibonacciNumber(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	//Achar o maior numero de Fibonacci menor que N
	f1, f2, f3 := 0, 1, 1

	for f3 <= n {
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}

	return f2
}

func representacaoFibonacci(n int) []int {
	valRepresentados := []int{}
	for n > 0 {

		f := closestFibonacciNumber(n)
		valRepresentados = append(valRepresentados, f)
		n = n - f
	}

	return valRepresentados
}

func Fibonacci(dados []byte) []string {

	result := []string{}
	Fibonacci := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

	for _, x := range dados {
		valRepresentados := representacaoFibonacci(int(x))
		output := ""
		j := len(valRepresentados) - 1

		for i := 0; j >= 0; i++ {
			if Fibonacci[i] == valRepresentados[j] {
				output = output + "1"
				j--
			} else {
				output = output + "0"
			}
		}
		output = output + "1"
		result = append(result, output)
	}
	return result
}
