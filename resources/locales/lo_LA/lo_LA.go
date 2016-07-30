package lo_LA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lo_LA struct {
	locale string
}

// New returns a new instance of translator for the 'lo_LA' locale
func New() locales.Translator {
	return &lo_LA{
		locale: "lo_LA",
	}
}

// Locale returns the current translators string locale
func (l *lo_LA) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lo_LA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
