package ig

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ig struct {
	locale string
}

// New returns a new instance of translator for the 'ig' locale
func New() locales.Translator {
	return &ig{
		locale: "ig",
	}
}

// Locale returns the current translators string locale
func (l *ig) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ig) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
