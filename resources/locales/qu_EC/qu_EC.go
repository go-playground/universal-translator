package qu_EC

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_EC struct {
	locale string
}

// New returns a new instance of translator for the 'qu_EC' locale
func New() locales.Translator {
	return &qu_EC{
		locale: "qu_EC",
	}
}

// Locale returns the current translators string locale
func (l *qu_EC) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *qu_EC) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
