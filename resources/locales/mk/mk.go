package mk

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mk struct {
	locale string
}

// New returns a new instance of translator for the 'mk' locale
func New() locales.Translator {
	return &mk{
		locale: "mk",
	}
}

// Locale returns the current translators string locale
func (l *mk) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mk) CardinalPluralRule(num string) (locales.PluralRule, error) {

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if (v == 0 && i%10 == 1) || (f%10 == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
