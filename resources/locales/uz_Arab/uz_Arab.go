package uz_Arab

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Arab struct {
	locale string
}

// New returns a new instance of translator for the 'uz_Arab' locale
func New() locales.Translator {
	return &uz_Arab{
		locale: "uz_Arab",
	}
}

// Locale returns the current translators string locale
func (l *uz_Arab) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *uz_Arab) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
