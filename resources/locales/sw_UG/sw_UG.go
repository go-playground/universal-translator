package sw_UG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw_UG struct {
	locale string
}

// New returns a new instance of translator for the 'sw_UG' locale
func New() locales.Translator {
	return &sw_UG{
		locale: "sw_UG",
	}
}

// Locale returns the current translators string locale
func (l *sw_UG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sw_UG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
