package ms_MY

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_MY struct {
	locale string
}

// New returns a new instance of translator for the 'ms_MY' locale
func New() locales.Translator {
	return &ms_MY{
		locale: "ms_MY",
	}
}

// Locale returns the current translators string locale
func (l *ms_MY) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ms_MY) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
