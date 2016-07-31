package kl

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"DKK": ut.Currency{Currency: "DKK", DisplayName: "danmarkimut koruuni", Symbol: "kr."}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norskit koruuni", Symbol: "Nkr"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svenskit koruuni", Symbol: "Skr"}}
