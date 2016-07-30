package sq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sq struct {
	locale string
}

// New returns a new instance of translator for the 'sq' locale
func New() locales.Translator {
	return &sq{
		locale: "sq",
	}
}

// Locale returns the current translators string locale
func (l *sq) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sq) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
