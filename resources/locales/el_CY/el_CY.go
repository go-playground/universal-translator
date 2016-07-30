package el_CY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type el_CY struct {
	locale string
}

// New returns a new instance of translator for the 'el_CY' locale
func New() locales.Translator {
	return &el_CY{
		locale: "el_CY",
	}
}

// Locale returns the current translators string locale
func (l *el_CY) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *el_CY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
