package lu_CD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lu_CD struct {
	locale string
}

// New returns a new instance of translator for the 'lu_CD' locale
func New() locales.Translator {
	return &lu_CD{
		locale: "lu_CD",
	}
}

// Locale returns the current translators string locale
func (l *lu_CD) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lu_CD) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
