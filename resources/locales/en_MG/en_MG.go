package en_MG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_MG struct {
	locale string
}

// New returns a new instance of translator for the 'en_MG' locale
func New() locales.Translator {
	return &en_MG{
		locale: "en_MG",
	}
}

// Locale returns the current translators string locale
func (l *en_MG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *en_MG) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
