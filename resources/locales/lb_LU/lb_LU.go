package lb_LU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lb_LU struct {
	locale string
}

// New returns a new instance of translator for the 'lb_LU' locale
func New() locales.Translator {
	return &lb_LU{
		locale: "lb_LU",
	}
}

// Locale returns the current translators string locale
func (l *lb_LU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lb_LU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
