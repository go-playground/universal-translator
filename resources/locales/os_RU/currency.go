package os_RU

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "RUB", DisplayName: "", Symbol: "₽"},
	}
}
