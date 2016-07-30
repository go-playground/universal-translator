package km

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type km struct {
	locale string
}

// New returns a new instance of translator for the 'km' locale
func New() locales.Translator {
	return &km{
		locale: "km",
	}
}

// Locale returns the current translators string locale
func (l *km) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *km) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
