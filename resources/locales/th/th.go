package th

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type th struct {
	locale string
}

// New returns a new instance of translator for the 'th' locale
func New() locales.Translator {
	return &th{
		locale: "th",
	}
}

// Locale returns the current translators string locale
func (l *th) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *th) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
