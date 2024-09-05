package main

import "github.com/charmbracelet/huh"

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
			huh.NewSelect[currency]().Title("Which Currency do you want to convert from?").Options(
				huh.NewOption("USD $", currency{"USD", "0"}),
				huh.NewOption("EUR €", currency{"EUR", "0"}),
				huh.NewOption("GBP £", currency{"GBP", "0"}),
			).Value(&startCurrency),
			huh.NewInput().Title("Enter an amount").Placeholder("0").Value(&startCurrency.value),
			huh.NewSelect[currency]().Title("Which Currency do you want to convert to?").Options(
				huh.NewOption("USD $", currency{"USD", "0"}),
				huh.NewOption("EUR €", currency{"EUR", "0"}),
				huh.NewOption("GBP £", currency{"GBP", "0"}),
			),
			huh.NewConfirm().Title("Are you sure you want to convert?").Value(&confirm),
		),
	)

	err := form.Run()
	if err != nil {
		panic(err)
	}

}
