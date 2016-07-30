package sbp_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sbp_TZ struct {
	locale string
}

// New returns a new instance of translator for the 'sbp_TZ' locale
func New() locales.Translator {
	return &sbp_TZ{
		locale: "sbp_TZ",
	}
}

// Locale returns the current translators string locale
func (l *sbp_TZ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sbp_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
