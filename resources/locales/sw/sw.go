package sw

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw struct {
	locale string
}

// New returns a new instance of translator for the 'sw' locale
func New() locales.Translator {
	return &sw{
		locale: "sw",
	}
}

// Locale returns the current translators string locale
func (l *sw) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sw) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
