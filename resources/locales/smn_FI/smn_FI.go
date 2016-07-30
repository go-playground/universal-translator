package smn_FI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type smn_FI struct {
	locale string
}

// New returns a new instance of translator for the 'smn_FI' locale
func New() locales.Translator {
	return &smn_FI{
		locale: "smn_FI",
	}
}

// Locale returns the current translators string locale
func (l *smn_FI) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *smn_FI) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	}

	return locales.PluralRuleOther, nil
}
