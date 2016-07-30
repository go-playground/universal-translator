package smn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type smn struct {
	locale string
}

// New returns a new instance of translator for the 'smn' locale
func New() locales.Translator {
	return &smn{
		locale: "smn",
	}
}

// Locale returns the current translators string locale
func (l *smn) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *smn) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
