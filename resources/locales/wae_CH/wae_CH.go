package wae_CH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type wae_CH struct {
	locale string
}

// New returns a new instance of translator for the 'wae_CH' locale
func New() locales.Translator {
	return &wae_CH{
		locale: "wae_CH",
	}
}

// Locale returns the current translators string locale
func (l *wae_CH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *wae_CH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
