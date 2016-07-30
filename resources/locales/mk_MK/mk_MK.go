package mk_MK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mk_MK struct {
	locale string
}

// New returns a new instance of translator for the 'mk_MK' locale
func New() locales.Translator {
	return &mk_MK{
		locale: "mk_MK",
	}
}

// Locale returns the current translators string locale
func (l *mk_MK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mk_MK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (v == 0 && i%10 == 1) || (f%10 == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
