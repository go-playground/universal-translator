package os_GE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type os_GE struct {
	locale string
}

// New returns a new instance of translator for the 'os_GE' locale
func New() locales.Translator {
	return &os_GE{
		locale: "os_GE",
	}
}

// Locale returns the current translators string locale
func (l *os_GE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *os_GE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
