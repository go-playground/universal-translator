package fa_AF

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AUD", DisplayName: "دالر آسترالیا", Symbol: ""},
		{Currency: "BND", DisplayName: "دالر برونی", Symbol: ""},
		{Currency: "BYR", DisplayName: "روبل روسیهٔ سفید", Symbol: ""},
		{Currency: "CAD", DisplayName: "دالر کانادا", Symbol: ""},
		{Currency: "CHF", DisplayName: "فرانک سویس", Symbol: ""},
		{Currency: "DKK", DisplayName: "کرون دنمارک", Symbol: ""},
		{Currency: "JPY", DisplayName: "ین جاپان", Symbol: ""},
		{Currency: "MXN", DisplayName: "پزوی مکسیکو", Symbol: ""},
		{Currency: "NLG", DisplayName: "گیلدر هالند", Symbol: ""},
		{Currency: "NOK", DisplayName: "کرون ناروی", Symbol: ""},
		{Currency: "SEK", DisplayName: "کرون سویدن", Symbol: ""},
		{Currency: "SGD", DisplayName: "دالر سینگاپور", Symbol: ""},
		{Currency: "TJS", DisplayName: "سامانی تاجکستان", Symbol: ""},
		{Currency: "USD", DisplayName: "دالر امریکا", Symbol: ""},
	}
}
