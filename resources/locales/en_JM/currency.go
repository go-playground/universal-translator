package en_JM

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "JMD", DisplayName: "", Symbol: "$"},
	}
}
