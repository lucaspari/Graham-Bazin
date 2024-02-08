package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	var option int16
	var valor string
	log.Println("Deseja fazer o calculo para Ações ou Fundos Imobiliários? (1/2)")
	_, err := fmt.Scan(&option)
	log.Println("Digite o valor do ticker")
	_, err = fmt.Scan(&valor)
	if err != nil {
		log.Fatal("Erro ao ler valor", err)
	}
	done := make(chan bool)
	go func() {
		visible := false
		for {
			select {
			case <-done:
				return
			default:
				if visible {
					fmt.Printf("\r               \r")
				} else {
					fmt.Printf("\rCarregando...")
				}
				visible = !visible
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	cotacao, dividendYield, lpa, vpa := getStockValue(valor, option)
	done <- true
	log.Println("Deseja utilizar o método Bazin ou Graham? (1/2)")
	_, err = fmt.Scan(&option)
	switch option {
	case 1:
		dividendYield, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(dividendYield, "%", ""), ",", "."), 64)
		if err != nil {
			log.Fatal(err)
		}
		cotacao, err := strconv.ParseFloat(strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(cotacao, "R$", ""), ",", ".")), 64)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Valor Teto: %.2f\n", bazinMethod(dividendYield, cotacao))

	case 2:
		vpa, _ := strconv.ParseFloat(strings.ReplaceAll(vpa, ",", "."), 64)
		lpa, _ := strconv.ParseFloat(strings.ReplaceAll(lpa, ",", "."), 64)
		log.Printf("Valor Intríseco: %.2f\n", grahamMethod(lpa, vpa))
	}
}
