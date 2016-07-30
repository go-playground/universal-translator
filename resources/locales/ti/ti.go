package ti

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ti struct {
	locale string
}

// New returns a new instance of translator for the 'ti' locale
func New() locales.Translator {
	return &ti{
		locale: "ti",
	}
}

// Locale returns the current translators string locale
func (l *ti) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ti) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
