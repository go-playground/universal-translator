package sg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sg struct {
	locale string
}

// New returns a new instance of translator for the 'sg' locale
func New() locales.Translator {
	return &sg{
		locale: "sg",
	}
}

// Locale returns the current translators string locale
func (l *sg) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sg) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
