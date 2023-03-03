// Compress files concurrently using a wait group.
package main

import (
	"log"
	"os"
	"sync"

	"github.com/jreisinger/gokatas/compress"
)

// Wg.Wait: é uma chamada bloqueante. Só vou começar a próxima leva de threads quando as threads que estão antes, terminarem
// RAIZA: ela disse que ele eh uma chamada bloqueante sim, mas não é que ele vai começar a próxima goroutine quando a antiga
// acabar, e sim que ele só vai terminar de rodar o que estiver na main (pra além da thread), quando a thread ou as threads
// acabarem. Motivo: quando uma thread começa a rodar, é como se ela fizesse isso num espaço separado. Se deixarmos todo o
// resto de código da main rodando enquanto a thread roda, sem um Wait Group inicializado, a main vai ser interamente
// executada antes de a thread acabar, e aí teremos perdido tudo o que a thread iria fazer.

func main() {
	var wg sync.WaitGroup
	for _, arg := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if err := compress.Compress(file); err != nil {
				log.Printf("compressing: %s %v", file, err)
			}
		}(arg)
	}
	wg.Done()
}
