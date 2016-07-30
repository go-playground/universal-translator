package ms

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms struct {
	locale string
}

// New returns a new instance of translator for the 'ms' locale
func New() locales.Translator {
	return &ms{
		locale: "ms",
	}
}

// Locale returns the current translators string locale
func (l *ms) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ms) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
