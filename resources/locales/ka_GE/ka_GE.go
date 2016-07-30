package ka_GE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ka_GE struct {
	locale string
}

// New returns a new instance of translator for the 'ka_GE' locale
func New() locales.Translator {
	return &ka_GE{
		locale: "ka_GE",
	}
}

// Locale returns the current translators string locale
func (l *ka_GE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ka_GE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
