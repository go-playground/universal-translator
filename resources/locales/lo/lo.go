package lo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lo struct {
	locale string
}

// New returns a new instance of translator for the 'lo' locale
func New() locales.Translator {
	return &lo{
		locale: "lo",
	}
}

// Locale returns the current translators string locale
func (l *lo) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lo) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
