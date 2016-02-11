package az_Cyrl

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AZN", DisplayName: "манат", Symbol: "₼"},
	}
}
