package hu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type hu struct {
	locale string
}

// New returns a new instance of translator for the 'hu' locale
func New() locales.Translator {
	return &hu{
		locale: "hu",
	}
}

// Locale returns the current translators string locale
func (l *hu) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *hu) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
