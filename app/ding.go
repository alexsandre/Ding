package app

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const urlDimpes = "http://dimpes.mpes.mp.br/ResultadoPesquisa"

func Pesquisar(termo string, dataInicial time.Time, dataFinal time.Time) (*[]Resultado, error) {
	if termo == "" {
		return nil, errors.New("Termo não pode ser vazio")
	}

	formData := montarFormData(termo, dataInicial, dataFinal)

	resp, err := http.PostForm(urlDimpes, formData)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("Erro ao obter o resultado da pesquisa")
	}

	htmlDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var retorno []Resultado
	htmlDoc.Find("tr .odd").Each(func(i int, s *goquery.Selection) {

		retorno = append(retorno, Resultado{
			Text:               "A",
			LinkPagina:         "B",
			LinkEdicaoCompleta: "C",
			LinkVisualizar:     "D",
		})
	})

	return nil, nil
}

func montarFormData(termo string, dataInicial time.Time, dataFinal time.Time) map[string][]string {
	return url.Values{
		"palavra": {termo},
		"dataDe":  {dataInicial.Format("02-01-2006")},
		"dataAte": {dataFinal.Format("02-01-2006")},
	}
}

type Resultado struct {
	Text               string
	LinkPagina         string
	LinkEdicaoCompleta string
	LinkVisualizar     string
}
