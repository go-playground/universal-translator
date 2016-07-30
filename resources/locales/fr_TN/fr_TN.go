package fr_TN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_TN struct {
	locale string
}

// New returns a new instance of translator for the 'fr_TN' locale
func New() locales.Translator {
	return &fr_TN{
		locale: "fr_TN",
	}
}

// Locale returns the current translators string locale
func (l *fr_TN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fr_TN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
