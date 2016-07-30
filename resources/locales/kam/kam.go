package kam

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kam struct {
	locale string
}

// New returns a new instance of translator for the 'kam' locale
func New() locales.Translator {
	return &kam{
		locale: "kam",
	}
}

// Locale returns the current translators string locale
func (l *kam) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kam) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
