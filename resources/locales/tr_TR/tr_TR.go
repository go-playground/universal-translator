package tr_TR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type tr_TR struct {
	locale string
}

// New returns a new instance of translator for the 'tr_TR' locale
func New() locales.Translator {
	return &tr_TR{
		locale: "tr_TR",
	}
}

// Locale returns the current translators string locale
func (l *tr_TR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *tr_TR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
