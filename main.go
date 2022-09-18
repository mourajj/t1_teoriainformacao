package main

import (
	"fmt"
	"teoriaInformacao/encoding"
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

func main() {

	//data, err := os.ReadFile("files/test.txt")
	//check(err)

	//fmt.Println(data)

	testData := []byte{2, 3, 0, 1, 1, 6, 2, 4, 4, 4, 1, 3, 5, 2}

	fmt.Println(encoding.Golomb(4, testData))
	//fmt.Println(encoding.Unario(testData))

}
