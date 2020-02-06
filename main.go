package ding

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Clipping(w http.ResponseWriter, r *http.Request) {
	doc, err := requestHtmlSearch()
	if err != nil {
		return
	}

	if (*doc).Find("tr.odd").Length() > 0 {
		fmt.Println("Deu certo")
	} else {
		fmt.Println("Não deu certo")
	}
}

func requestHtmlSearch() (*goquery.Document, error) {
	urlDimpes := "http://dimpes.mpes.mp.br/ResultadoPesquisa"
	termoBusca := os.Getenv("termoBusca")

	formData := url.Values{
		"palavra": {termoBusca},
		"dataDe":  {time.Now().Local().Format("02-01-2006")},
		"dataAte": {time.Now().Local().Format("02-01-2006")},
	}

	resp, err := http.PostForm(urlDimpes, formData)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("StatusCode != 200")
	}

	htmlDoc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return htmlDoc, nil
}
