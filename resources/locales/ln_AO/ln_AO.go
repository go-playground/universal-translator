package ln_AO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ln_AO struct {
	locale string
}

// New returns a new instance of translator for the 'ln_AO' locale
func New() locales.Translator {
	return &ln_AO{
		locale: "ln_AO",
	}
}

// Locale returns the current translators string locale
func (l *ln_AO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ln_AO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
