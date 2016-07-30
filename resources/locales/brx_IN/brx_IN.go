package brx_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type brx_IN struct {
	locale string
}

// New returns a new instance of translator for the 'brx_IN' locale
func New() locales.Translator {
	return &brx_IN{
		locale: "brx_IN",
	}
}

// Locale returns the current translators string locale
func (l *brx_IN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *brx_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
