package bg_BG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bg_BG struct {
	locale string
}

// New returns a new instance of translator for the 'bg_BG' locale
func New() locales.Translator {
	return &bg_BG{
		locale: "bg_BG",
	}
}

// Locale returns the current translators string locale
func (l *bg_BG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bg_BG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
