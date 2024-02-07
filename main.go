package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

func getStockValue(valor string) {
	response, err := http.Get("https://investidor10.com.br/acoes/" + valor + "/")
	if err != nil {
		log.Fatal("Erro ao fazer requisição", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if response.StatusCode != 200 {
		log.Fatal("Erro ao fazer requisição", err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Erro ao criar documento", err)
	}
	doc.Find("div._card-body > div > span").Each(func(index int, item *goquery.Selection) {
		texts := item.Text()
		if strings.Contains(texts, "R$") {
			log.Println(valor, ": ", texts)
		}
	})
	if err != nil {
		log.Fatal("Erro ao fazer requisição", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if response.StatusCode != 200 {
		log.Fatal("Erro ao fazer requisição", err)
	}
	doc, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Erro ao criar documento", err)
	}
	doc.Find("div._card-body > div > span").Each(func(index int, item *goquery.Selection) {
		texts := item.Text()
		if strings.Contains(texts, "R$") {
			log.Println(": ", texts)
		}
	})
}
func main() {
	log.Println("Digite o valor do ticker")
	var valor string
	_, err := fmt.Scan(&valor)
	if err != nil {
		log.Fatal("Erro ao ler valor", err)
	}
	getStockValue(valor)
}
