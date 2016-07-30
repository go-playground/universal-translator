package si_LK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type si_LK struct {
	locale string
}

// New returns a new instance of translator for the 'si_LK' locale
func New() locales.Translator {
	return &si_LK{
		locale: "si_LK",
	}
}

// Locale returns the current translators string locale
func (l *si_LK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *si_LK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n == 0 || n == 1) || (i == 0 && f == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
