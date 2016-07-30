package kab

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kab struct {
	locale string
}

// New returns a new instance of translator for the 'kab' locale
func New() locales.Translator {
	return &kab{
		locale: "kab",
	}
}

// Locale returns the current translators string locale
func (l *kab) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kab) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
