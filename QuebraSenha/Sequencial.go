package main

import (
    "fmt" //pacote da biblioteca padrao do go
	"time"
)

func main() {
    senhaAlvo:= "12345678"
	inicio:= time.Now()

	for i:= 0; i<=99999999; i++ {
		tentativa:= fmt.Sprintf("%08d", i) //%d quer dizer numero inteiro e 08 que vai ser de 8 digitos
		if tentativa == senhaAlvo {
			fmt.Println("Senha Encontrada: ", tentativa)
			break
		}

	}
	duracao:= time.Since(inicio)
	fmt.Println("Tempo: ", duracao)
}
