package fr_BJ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_BJ struct {
	locale string
}

// New returns a new instance of translator for the 'fr_BJ' locale
func New() locales.Translator {
	return &fr_BJ{
		locale: "fr_BJ",
	}
}

// Locale returns the current translators string locale
func (l *fr_BJ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fr_BJ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
