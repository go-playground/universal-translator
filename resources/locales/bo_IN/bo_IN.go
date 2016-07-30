package bo_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bo_IN struct {
	locale string
}

// New returns a new instance of translator for the 'bo_IN' locale
func New() locales.Translator {
	return &bo_IN{
		locale: "bo_IN",
	}
}

// Locale returns the current translators string locale
func (l *bo_IN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bo_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
