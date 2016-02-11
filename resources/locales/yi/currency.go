package yi

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "בראזיל רעאל", Symbol: ""},
		{Currency: "BZD", DisplayName: "בעליז דאלאַר", Symbol: ""},
		{Currency: "CHF", DisplayName: "שווייצער פֿראַנק", Symbol: ""},
		{Currency: "CNY", DisplayName: "כינעזישער יואן", Symbol: ""},
		{Currency: "EUR", DisplayName: "איירא", Symbol: "€"},
		{Currency: "GBP", DisplayName: "פֿונט שטערלינג", Symbol: "£"},
		{Currency: "INR", DisplayName: "אינדישער רופי", Symbol: "₹"},
		{Currency: "JPY", DisplayName: "יאפאנעזישער יען", Symbol: "JP¥"},
		{Currency: "RUB", DisplayName: "רוסישער רובל", Symbol: "RUB"},
		{Currency: "SEK", DisplayName: "שוועדישע קראנע", Symbol: ""},
		{Currency: "USD", DisplayName: "אמעריקאנער דאלאר", Symbol: "$"},
		{Currency: "XAG", DisplayName: "זילבער", Symbol: ""},
		{Currency: "XAU", DisplayName: "גאלד", Symbol: ""},
		{Currency: "XXX", DisplayName: "אומבאַוואוסטע וואַלוטע", Symbol: ""},
	}
}
