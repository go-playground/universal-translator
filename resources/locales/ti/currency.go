package ti

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "የብራዚል ሪል", Symbol: ""},
		{Currency: "CNY", DisplayName: "የቻይና ዩአን ረንሚንቢ", Symbol: ""},
		{Currency: "ETB", DisplayName: "የኢትዮጵያ ብር", Symbol: "Br"},
		{Currency: "EUR", DisplayName: "አውሮ", Symbol: ""},
		{Currency: "GBP", DisplayName: "የእንግሊዝ ፓውንድ ስተርሊንግ", Symbol: ""},
		{Currency: "INR", DisplayName: "የሕንድ ሩፒ", Symbol: ""},
		{Currency: "JPY", DisplayName: "የጃፓን የን", Symbol: ""},
		{Currency: "RUB", DisplayName: "የራሻ ሩብል", Symbol: ""},
		{Currency: "USD", DisplayName: "የአሜሪካን ዶላር", Symbol: ""},
	}
}
