package ur_IN

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "CRC", DisplayName: "کوسٹا ریکا کولون", Symbol: ""},
		{Currency: "CUC", DisplayName: "قابل منتقلی کیوبائی پیسو", Symbol: ""},
		{Currency: "CUP", DisplayName: "کیوبائی پیسو", Symbol: ""},
		{Currency: "CVE", DisplayName: "کیپ ورڈی اسکیوڈو", Symbol: ""},
		{Currency: "ERN", DisplayName: "اریٹیریائی ناکفا", Symbol: ""},
		{Currency: "GBP", DisplayName: "برطانوی پاونڈ سٹرلنگ", Symbol: ""},
		{Currency: "GHS", DisplayName: "گھانی سیڈی", Symbol: ""},
		{Currency: "PKR", DisplayName: "پاکستانی روپیہ", Symbol: "PKR"},
		{Currency: "WST", DisplayName: "ساموآئی ٹالا", Symbol: ""},
	}
}
