package sbp

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sbp struct {
	locale string
}

// New returns a new instance of translator for the 'sbp' locale
func New() locales.Translator {
	return &sbp{
		locale: "sbp",
	}
}

// Locale returns the current translators string locale
func (l *sbp) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sbp) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
