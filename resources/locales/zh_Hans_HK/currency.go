package zh_Hans_HK

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AWG", DisplayName: "阿鲁巴弗罗林", Symbol: ""},
		{Currency: "CNY", DisplayName: "", Symbol: "CN¥"},
		{Currency: "KYD", DisplayName: "开曼群岛元", Symbol: ""},
		{Currency: "NIO", DisplayName: "尼加拉瓜科多巴", Symbol: ""},
		{Currency: "XAG", DisplayName: "白银", Symbol: ""},
	}
}
