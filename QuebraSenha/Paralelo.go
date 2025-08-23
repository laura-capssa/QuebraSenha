package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    senhaAlvo := "12345678"
    N := 4 // numero de threads

    var wg sync.WaitGroup
    found := make(chan string, 1) // canal bufferizado pra avisar

    inicio := time.Now()

    for t := 0; t < N; t++ {
        ini := t * (100000000 / N)
        fim := (t+1) * (100000000 / N)

        wg.Add(1)
        go func(inicio, fim int) {
            defer wg.Done()
            for i := inicio; i < fim; i++ {
                tentativa := fmt.Sprintf("%08d", i)
                if tentativa == senhaAlvo {
                    found <- tentativa 
                    return
                }
            }
        }(ini, fim)
    }

    // espera ou pela senha ou pelo fim das threads
    go func() {
        wg.Wait()
        close(found)
    }()

    if senha := <-found; senha != "" {
        fmt.Println("Achei a senha:", senha)
    }

    duracao := time.Since(inicio)
    fmt.Println("Tempo:", duracao)
}
