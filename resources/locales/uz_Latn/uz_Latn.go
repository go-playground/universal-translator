package uz_Latn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type uz_Latn struct {
	locale string
}

// New returns a new instance of translator for the 'uz_Latn' locale
func New() locales.Translator {
	return &uz_Latn{
		locale: "uz_Latn",
	}
}

// Locale returns the current translators string locale
func (l *uz_Latn) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *uz_Latn) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
