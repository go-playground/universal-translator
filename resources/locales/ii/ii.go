package ii

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ii struct {
	locale string
}

// New returns a new instance of translator for the 'ii' locale
func New() locales.Translator {
	return &ii{
		locale: "ii",
	}
}

// Locale returns the current translators string locale
func (l *ii) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ii) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
