package kde_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kde_TZ struct {
	locale string
}

// New returns a new instance of translator for the 'kde_TZ' locale
func New() locales.Translator {
	return &kde_TZ{
		locale: "kde_TZ",
	}
}

// Locale returns the current translators string locale
func (l *kde_TZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kde_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
