package ha

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ha struct {
	locale string
}

// New returns a new instance of translator for the 'ha' locale
func New() locales.Translator {
	return &ha{
		locale: "ha",
	}
}

// Locale returns the current translators string locale
func (l *ha) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ha) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
