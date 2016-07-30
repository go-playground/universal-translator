package ki

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ki struct {
	locale string
}

// New returns a new instance of translator for the 'ki' locale
func New() locales.Translator {
	return &ki{
		locale: "ki",
	}
}

// Locale returns the current translators string locale
func (l *ki) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ki) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
