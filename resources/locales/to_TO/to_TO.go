package to_TO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type to_TO struct {
	locale string
}

// New returns a new instance of translator for the 'to_TO' locale
func New() locales.Translator {
	return &to_TO{
		locale: "to_TO",
	}
}

// Locale returns the current translators string locale
func (l *to_TO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *to_TO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
