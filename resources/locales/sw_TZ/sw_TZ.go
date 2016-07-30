package sw_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw_TZ struct {
	locale string
}

// New returns a new instance of translator for the 'sw_TZ' locale
func New() locales.Translator {
	return &sw_TZ{
		locale: "sw_TZ",
	}
}

// Locale returns the current translators string locale
func (l *sw_TZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sw_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
