package root

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type root struct {
	locale string
}

// New returns a new instance of translator for the 'root' locale
func New() locales.Translator {
	return &root{
		locale: "root",
	}
}

// Locale returns the current translators string locale
func (l *root) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *root) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
