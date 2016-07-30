package id_ID

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type id_ID struct {
	locale string
}

// New returns a new instance of translator for the 'id_ID' locale
func New() locales.Translator {
	return &id_ID{
		locale: "id_ID",
	}
}

// Locale returns the current translators string locale
func (l *id_ID) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *id_ID) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
