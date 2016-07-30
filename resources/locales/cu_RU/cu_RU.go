package cu_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cu_RU struct {
	locale string
}

// New returns a new instance of translator for the 'cu_RU' locale
func New() locales.Translator {
	return &cu_RU{
		locale: "cu_RU",
	}
}

// Locale returns the current translators string locale
func (l *cu_RU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *cu_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
