package aws

import (
	"os"
	"time"

	"github.com/alexsandre/ding/app"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Handler(ProcurarTermo)
}

func ProcurarTermo() (string, error) {
	termo := os.Getenv("termo")
	dataInicial := time.Now().UTC()
	dataFinal := time.Now().UTC()

	retorno, err := app.Pesquisar(termo, dataInicial, dataFinal)
	if err != nil {
		return "", err
	}
}
