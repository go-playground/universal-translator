package cu

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "BRL", DisplayName: "бразі́льскїй реа́лъ", Symbol: "R$"},
		{Currency: "BYR", DisplayName: "бѣлорꙋ́сскїй рꙋ́бль", Symbol: "BYR"},
		{Currency: "CNY", DisplayName: "хи́нскїй ю҆а́нь", Symbol: "CN¥"},
		{Currency: "EUR", DisplayName: "є҆́ѵрѡ", Symbol: "€"},
		{Currency: "GBP", DisplayName: "а҆нглі́йскїй фꙋ́нтъ сте́рлингѡвъ", Symbol: "£"},
		{Currency: "INR", DisplayName: "і҆нді́йскаѧ рꙋ́пїѧ", Symbol: "₹"},
		{Currency: "JPY", DisplayName: "ꙗ҆пѡ́нскаѧ і҆е́на", Symbol: "JP¥"},
		{Currency: "KGS", DisplayName: "кирги́зскїй сꙋ́мъ", Symbol: "KGS"},
		{Currency: "KZT", DisplayName: "каза́хскаѧ деньга̀", Symbol: "₸"},
		{Currency: "RUB", DisplayName: "рѡссі́йскїй рꙋ́бль", Symbol: "₽"},
		{Currency: "UAH", DisplayName: "\u1c82у҆краи́нскаѧ гри́вна", Symbol: "₴"},
		{Currency: "USD", DisplayName: "а҆мерїка́нскїй до́лларъ", Symbol: "$"},
		{Currency: "XXX", DisplayName: "невѣ́домое пла́тное сре́дство", Symbol: ""},
	}
}
