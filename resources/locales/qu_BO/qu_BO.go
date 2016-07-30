package qu_BO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type qu_BO struct {
	locale string
}

// New returns a new instance of translator for the 'qu_BO' locale
func New() locales.Translator {
	return &qu_BO{
		locale: "qu_BO",
	}
}

// Locale returns the current translators string locale
func (l *qu_BO) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *qu_BO) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
