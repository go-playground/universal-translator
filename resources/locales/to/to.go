package to

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type to struct {
	locale string
}

// New returns a new instance of translator for the 'to' locale
func New() locales.Translator {
	return &to{
		locale: "to",
	}
}

// Locale returns the current translators string locale
func (l *to) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *to) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
