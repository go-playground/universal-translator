package bem_ZM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bem_ZM struct {
	locale string
}

// New returns a new instance of translator for the 'bem_ZM' locale
func New() locales.Translator {
	return &bem_ZM{
		locale: "bem_ZM",
	}
}

// Locale returns the current translators string locale
func (l *bem_ZM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bem_ZM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
