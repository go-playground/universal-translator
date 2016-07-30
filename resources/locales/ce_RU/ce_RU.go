package ce_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ce_RU struct {
	locale string
}

// New returns a new instance of translator for the 'ce_RU' locale
func New() locales.Translator {
	return &ce_RU{
		locale: "ce_RU",
	}
}

// Locale returns the current translators string locale
func (l *ce_RU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ce_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
