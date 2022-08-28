package golomb

import (
	"fmt"
	"math"
)

//https://en.wikipedia.org/wiki/Golomb_coding
//http://multimedia.ufp.pt/codecs/compressao-sem-perdas/codificacao-estatistica/codificacao-de-golomb-rice/

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
