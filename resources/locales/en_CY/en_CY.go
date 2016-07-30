package en_CY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_CY struct {
	locale string
}

// New returns a new instance of translator for the 'en_CY' locale
func New() locales.Translator {
	return &en_CY{
		locale: "en_CY",
	}
}

// Locale returns the current translators string locale
func (l *en_CY) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *en_CY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
