package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type currency struct {
	name  string
	value string
}

var resultCurrency currency

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Enter an amount").Placeholder("0").Value(&resultCurrency.value),
			huh.NewSelect[string]().Title("Which Currency do you want to convert to?").Options(
				huh.NewOption("EUR €", "EUR"),
				huh.NewOption("GBP £", "GBP"),
			).Value(&resultCurrency.name),
		),
	)

	err := form.Run()
	if err != nil {
		panic(err)
	}

	convert := func() {
		resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=" + os.Getenv("CURRENCY_API_KEY"))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var data map[string]interface{}
		json.Unmarshal(body, &data)

		rates := data["rates"].(map[string]interface{})
		for key, value := range rates {
			if key == resultCurrency.name {
				resultFloat, err := strconv.ParseFloat(resultCurrency.value, 64)
				if err != nil {
					panic(err)
				}
				resultCurrency.value = fmt.Sprintf("%.2f", resultFloat*value.(float64))
			}
		}
	}

	err = spinner.New().Title("Converting currency...").Action(convert).Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Converted Amount: %s %s\n", resultCurrency.value, resultCurrency.name)
}
