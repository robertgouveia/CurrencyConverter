package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type currency struct {
	name  string
	value string
}

var (
	startCurrency  currency
	resultCurrency currency
	confirm        bool
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Which Currency do you want to convert from?").Options(
				huh.NewOption("USD $", "USD"),
				huh.NewOption("EUR €", "EUR"),
				huh.NewOption("GBP £", "GBP"),
			).Value(&startCurrency.name),
			huh.NewInput().Title("Enter an amount").Placeholder("0").Value(&startCurrency.value),
			huh.NewSelect[string]().Title("Which Currency do you want to convert to?").Options(
				huh.NewOption("USD $", "USD"),
				huh.NewOption("EUR €", "EUR"),
				huh.NewOption("GBP £", "GBP"),
			).Value(&resultCurrency.name),
			huh.NewConfirm().Title("Are you sure you want to convert?").Value(&confirm),
		),
	)

	err := form.Run()
	if err != nil {
		panic(err)
	}

	convert := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Converting...")
		resultCurrency.value = "100"
	}

	err = spinner.New().Title("Converting currency...").Action(convert).Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(resultCurrency.name, ":", resultCurrency.value)
}
