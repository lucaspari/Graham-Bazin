package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
)

func getStockValue(valor string, opcao int16) (string, string, string, string) {
	var response *http.Response
	var err error
	if opcao == 1 {
		response, err = http.Get("https://investidor10.com.br/acoes/" + valor + "/")
	} else {
		response, err = http.Get("https://investidor10.com.br/fiis/" + valor + "/")
	}
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
	cotacao := findCotacao(doc)
	dividendYield := findDividendYield(doc, opcao)
	lucroporAcao := findLucroPorAcao(doc)
	valorPatrimonialPorAcao := findValorPorAcao(doc)
	return cotacao, dividendYield, lucroporAcao, valorPatrimonialPorAcao
}

func findDividendYield(doc *goquery.Document, opcao int16) string {
	var result string
	if opcao == 1 {
		doc.Find("#cards-ticker > div._card.dy > div._card-body > span").Each(func(index int, item *goquery.Selection) {
			log.Println("Dividend Yield: ", strings.TrimSpace(item.Text()))
			result = strings.TrimSpace(item.Text())
		})
	} else {
		doc.Find("#cards-ticker > div:nth-child(2) > div._card-body > div > span").Each(func(index int, item *goquery.Selection) {
			log.Println("Dividend Yield: ", strings.TrimSpace(item.Text()))
			result = strings.TrimSpace(item.Text())
		})
	}
	return result
}
func findCotacao(doc *goquery.Document) string {
	var result string
	doc.Find("#cards-ticker > div._card.cotacao > div._card-body > div > span").Each(func(index int, item *goquery.Selection) {
		log.Println("Cotação: ", strings.TrimSpace(item.Text()))
		result = strings.TrimSpace(item.Text())
	})
	return result
}

func findLucroPorAcao(doc *goquery.Document) string {
	var result string
	doc.Find("#table-indicators > div:nth-child(16) > div.value.d-flex.justify-content-between.align-items-center").Each(func(index int, item *goquery.Selection) {
		log.Println("Lucro por Ação (LPA): ", strings.TrimSpace(item.Text()))
		result = strings.TrimSpace(item.Text())
	})
	return strings.TrimSpace(result)
}

func findValorPorAcao(doc *goquery.Document) string {
	var result string
	doc.Find("#table-indicators > div:nth-child(15) > div.value.d-flex.justify-content-between.align-items-center").Each(func(index int, item *goquery.Selection) {
		log.Println("Valor Patrimonial por ação (VPA): ", strings.TrimSpace(item.Text()))
		result = strings.TrimSpace(item.Text())
	})
	return strings.TrimSpace(result)

}
func bazinMethod(dividendYield float64, cotacao float64) float64 {
	const BAZIN_VALUE = 0.06
	dividendosPagos := (dividendYield * cotacao) / 100
	return dividendosPagos / BAZIN_VALUE
}
func grahamMethod(lpa, vpa float64) float64 {
	const GRAHAM_VALUE = 22.5
	return math.Sqrt(GRAHAM_VALUE * (lpa * vpa))
}
