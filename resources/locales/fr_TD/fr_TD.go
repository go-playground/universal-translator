package fr_TD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fr_TD struct {
	locale string
}

// New returns a new instance of translator for the 'fr_TD' locale
func New() locales.Translator {
	return &fr_TD{
		locale: "fr_TD",
	}
}

// Locale returns the current translators string locale
func (l *fr_TD) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fr_TD) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
