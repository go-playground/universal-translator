package qu_PE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_PE struct {
	locale string
}

// New returns a new instance of translator for the 'qu_PE' locale
func New() locales.Translator {
	return &qu_PE{
		locale: "qu_PE",
	}
}

// Locale returns the current translators string locale
func (l *qu_PE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *qu_PE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
