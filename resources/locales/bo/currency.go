package bo

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CNY", DisplayName: "ཡུ་ཨན་", Symbol: "¥"},
		{Currency: "INR", DisplayName: "རྒྱ་གར་སྒོར་", Symbol: ""},
		{Currency: "USD", DisplayName: "ཨ་རིའི་སྒོར་", Symbol: ""},
		{Currency: "XXX", DisplayName: "མ་རྟོགས་པའི་ནུས་མེད་དངུལ་ལོར", Symbol: ""},
	}
}
