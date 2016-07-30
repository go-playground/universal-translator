package naq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type naq struct {
	locale string
}

// New returns a new instance of translator for the 'naq' locale
func New() locales.Translator {
	return &naq{
		locale: "naq",
	}
}

// Locale returns the current translators string locale
func (l *naq) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *naq) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	}

	return locales.PluralRuleOther, nil
}
