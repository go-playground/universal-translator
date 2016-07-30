package hy_AM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type hy_AM struct {
	locale string
}

// New returns a new instance of translator for the 'hy_AM' locale
func New() locales.Translator {
	return &hy_AM{
		locale: "hy_AM",
	}
}

// Locale returns the current translators string locale
func (l *hy_AM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *hy_AM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
