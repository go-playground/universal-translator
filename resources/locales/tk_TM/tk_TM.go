package tk_TM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type tk_TM struct {
	locale string
}

// New returns a new instance of translator for the 'tk_TM' locale
func New() locales.Translator {
	return &tk_TM{
		locale: "tk_TM",
	}
}

// Locale returns the current translators string locale
func (l *tk_TM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *tk_TM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
