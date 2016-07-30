package kl_GL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kl_GL struct {
	locale string
}

// New returns a new instance of translator for the 'kl_GL' locale
func New() locales.Translator {
	return &kl_GL{
		locale: "kl_GL",
	}
}

// Locale returns the current translators string locale
func (l *kl_GL) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kl_GL) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
