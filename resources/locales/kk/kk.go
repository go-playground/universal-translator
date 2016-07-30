package kk

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kk struct {
	locale string
}

// New returns a new instance of translator for the 'kk' locale
func New() locales.Translator {
	return &kk{
		locale: "kk",
	}
}

// Locale returns the current translators string locale
func (l *kk) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kk) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
