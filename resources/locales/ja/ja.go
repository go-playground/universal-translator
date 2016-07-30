package ja

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ja struct {
	locale string
}

// New returns a new instance of translator for the 'ja' locale
func New() locales.Translator {
	return &ja{
		locale: "ja",
	}
}

// Locale returns the current translators string locale
func (l *ja) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ja) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
