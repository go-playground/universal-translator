package br

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type br struct {
	locale string
}

// New returns a new instance of translator for the 'br' locale
func New() locales.Translator {
	return &br{
		locale: "br",
	}
}

// Locale returns the current translators string locale
func (l *br) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *br) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n%10 == 1 && (n%100 != 11 && n%100 != 71 && n%100 != 91) {
		return locales.PluralRuleOne, nil
	} else if n%10 == 2 && (n%100 != 12 && n%100 != 72 && n%100 != 92) {
		return locales.PluralRuleTwo, nil
	} else if n%10 >= 3 && n%10 <= 4 && (n%10 == 9) && (n%100 < 10 && n%100 > 19) || (n%100 < 70 && n%100 > 79) || (n%100 < 90 && n%100 > 99) {
		return locales.PluralRuleFew, nil
	} else if n != 0 && n%1000000 == 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
