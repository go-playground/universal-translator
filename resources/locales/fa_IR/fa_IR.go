package fa_IR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fa_IR struct {
	locale string
}

// New returns a new instance of translator for the 'fa_IR' locale
func New() locales.Translator {
	return &fa_IR{
		locale: "fa_IR",
	}
}

// Locale returns the current translators string locale
func (l *fa_IR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fa_IR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
