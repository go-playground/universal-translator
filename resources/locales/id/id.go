package id

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type id struct {
	locale string
}

// New returns a new instance of translator for the 'id' locale
func New() locales.Translator {
	return &id{
		locale: "id",
	}
}

// Locale returns the current translators string locale
func (l *id) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *id) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
