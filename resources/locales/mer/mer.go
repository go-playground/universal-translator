package mer

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mer struct {
	locale string
}

// New returns a new instance of translator for the 'mer' locale
func New() locales.Translator {
	return &mer{
		locale: "mer",
	}
}

// Locale returns the current translators string locale
func (l *mer) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mer) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
