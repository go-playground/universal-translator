package ta_MY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ta_MY struct {
	locale string
}

// New returns a new instance of translator for the 'ta_MY' locale
func New() locales.Translator {
	return &ta_MY{
		locale: "ta_MY",
	}
}

// Locale returns the current translators string locale
func (l *ta_MY) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ta_MY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
