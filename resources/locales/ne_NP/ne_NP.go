package ne_NP

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ne_NP struct {
	locale string
}

// New returns a new instance of translator for the 'ne_NP' locale
func New() locales.Translator {
	return &ne_NP{
		locale: "ne_NP",
	}
}

// Locale returns the current translators string locale
func (l *ne_NP) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ne_NP) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
