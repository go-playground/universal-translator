package en_FJ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_FJ struct {
	locale string
}

// New returns a new instance of translator for the 'en_FJ' locale
func New() locales.Translator {
	return &en_FJ{
		locale: "en_FJ",
	}
}

// Locale returns the current translators string locale
func (l *en_FJ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *en_FJ) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
