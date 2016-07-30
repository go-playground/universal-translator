package ha_GH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ha_GH struct {
	locale string
}

// New returns a new instance of translator for the 'ha_GH' locale
func New() locales.Translator {
	return &ha_GH{
		locale: "ha_GH",
	}
}

// Locale returns the current translators string locale
func (l *ha_GH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ha_GH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
