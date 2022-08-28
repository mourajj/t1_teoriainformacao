package golomb

import (
	"encoding/binary"
	"fmt"
	"math"
)

func golomb(divisor int, dados []byte) {

	if (divisor & (divisor - 1)) != 0 {
		panic(fmt.Sprintf("o divisor nao é uma potencia de 2"))
	}

	for _, x := range dados {
		value := binary.BigEndian.Uint64(x)
		q := math.Floor(divisor / value)
	}

}
