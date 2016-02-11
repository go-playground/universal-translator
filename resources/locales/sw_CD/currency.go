package sw_CD

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CDF", DisplayName: "", Symbol: "FC"},
		{Currency: "CNY", DisplayName: "Yuan Renminbi ya China", Symbol: ""},
		{Currency: "ETB", DisplayName: "Bir ya Uhabeshi", Symbol: ""},
		{Currency: "JPY", DisplayName: "Sarafu ya Kijapani", Symbol: ""},
		{Currency: "MGA", DisplayName: "Ariary ya Bukini", Symbol: ""},
		{Currency: "MRO", DisplayName: "Ugwiya ya Moritania", Symbol: ""},
		{Currency: "SCR", DisplayName: "Rupia ya Shelisheli", Symbol: ""},
		{Currency: "XAF", DisplayName: "Faranga CFA BEAC", Symbol: ""},
		{Currency: "XOF", DisplayName: "Faranga CFA BCEAO", Symbol: ""},
	}
}
