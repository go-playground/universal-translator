package mg_MG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mg_MG struct {
	locale string
}

// New returns a new instance of translator for the 'mg_MG' locale
func New() locales.Translator {
	return &mg_MG{
		locale: "mg_MG",
	}
}

// Locale returns the current translators string locale
func (l *mg_MG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mg_MG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
