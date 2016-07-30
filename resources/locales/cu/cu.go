package cu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cu struct {
	locale string
}

// New returns a new instance of translator for the 'cu' locale
func New() locales.Translator {
	return &cu{
		locale: "cu",
	}
}

// Locale returns the current translators string locale
func (l *cu) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *cu) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
