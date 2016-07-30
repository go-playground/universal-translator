package rm_CH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type rm_CH struct {
	locale string
}

// New returns a new instance of translator for the 'rm_CH' locale
func New() locales.Translator {
	return &rm_CH{
		locale: "rm_CH",
	}
}

// Locale returns the current translators string locale
func (l *rm_CH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *rm_CH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
