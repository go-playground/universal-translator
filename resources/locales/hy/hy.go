package hy

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type hy struct {
	locale string
}

// New returns a new instance of translator for the 'hy' locale
func New() locales.Translator {
	return &hy{
		locale: "hy",
	}
}

// Locale returns the current translators string locale
func (l *hy) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *hy) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
