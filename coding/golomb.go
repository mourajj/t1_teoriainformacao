package golomb

import (
	"fmt"
	"math"
)

//https://en.wikipedia.org/wiki/Golomb_coding
//http://multimedia.ufp.pt/codecs/compressao-sem-perdas/codificacao-estatistica/codificacao-de-golomb-rice/

//Passo a Passo Golomb-rice

//Converter o dígito para hexadecimal
//Dividir o valor hexadecimal pelo divisor selecionado gerando Y = Resto da Divisão, X = Resultado da Divisão
//Gerar o valor codificado X zeros, stop bit(1), Y representado em binário pela quantidade de bits resultante de logaritmo(divisor)

func golomb(divisor int, dados []byte) {

	if (divisor & (divisor - 1)) != 0 {
		panic(fmt.Sprintf("o divisor nao é uma potencia de 2"))
	}

	for _, x := range dados {
		value := int(x)
		q := math.Floor(float64(divisor / value))
		r := int(x) % divisor

	}

}
