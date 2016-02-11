package to

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AUD", DisplayName: "", Symbol: "AUD$"},
		{Currency: "BRL", DisplayName: "BRL", Symbol: ""},
		{Currency: "FJD", DisplayName: "", Symbol: "FJD"},
		{Currency: "GEL", DisplayName: "", Symbol: "₾"},
		{Currency: "NZD", DisplayName: "", Symbol: "NZD$"},
		{Currency: "PGK", DisplayName: "", Symbol: "PGK"},
		{Currency: "SBD", DisplayName: "", Symbol: "SBD"},
		{Currency: "TOP", DisplayName: "Paʻanga fakatonga", Symbol: "T$"},
		{Currency: "TWD", DisplayName: "", Symbol: "$"},
		{Currency: "VUV", DisplayName: "", Symbol: "VUV"},
		{Currency: "WST", DisplayName: "Tala fakahaʻamoa", Symbol: "WST"},
		{Currency: "XPF", DisplayName: "", Symbol: "CFPF"},
		{Currency: "XXX", DisplayName: "Pa’anga Ta’e’ilo", Symbol: ""},
	}
}
