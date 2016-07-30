package nnh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nnh struct {
	locale string
}

// New returns a new instance of translator for the 'nnh' locale
func New() locales.Translator {
	return &nnh{
		locale: "nnh",
	}
}

// Locale returns the current translators string locale
func (l *nnh) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nnh) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
