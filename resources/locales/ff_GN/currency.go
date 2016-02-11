package ff_GN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "GNF", DisplayName: "", Symbol: "FG"},
	}
}
