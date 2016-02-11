package mas_TZ

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "TZS", DisplayName: "", Symbol: "TSh"},
	}
}
