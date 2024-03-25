package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	escolhido := flag.Int("c", 1, "Número escolhido")
	maximo := flag.Int("m", 1, "Número máximo pra ser sorteado")
	flag.Parse()

	if *escolhido > *maximo {
		fmt.Println("O número escolhido não pode ser maior que o número máximo de sorteio.")
		os.Exit(1)
	}

	delayStr := os.Getenv("DELAY_SORTEIO")
	if delayStr == "" {
		fmt.Println("A variável DELAY_SORTEIO não está definida. Usando delay padrão de 5 segundos.")
		delayStr = "5"
	}

	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sorteando...")
	time.Sleep(time.Duration(delay) * time.Second)

	numeroSorteado := rand.Intn(*maximo-1) + 1
	if numeroSorteado == *escolhido {
		fmt.Printf("O valor sorteado foi %d e você escolheu %d. Parabéns !!!\n", numeroSorteado, *escolhido)
	} else {
		fmt.Printf("O valor sorteado foi %d e você escolheu %d. Tenta de novo...\n", numeroSorteado, *escolhido)
		os.Exit(1)
	}
}
