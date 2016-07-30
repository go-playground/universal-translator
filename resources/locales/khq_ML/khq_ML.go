package khq_ML

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type khq_ML struct {
	locale string
}

// New returns a new instance of translator for the 'khq_ML' locale
func New() locales.Translator {
	return &khq_ML{
		locale: "khq_ML",
	}
}

// Locale returns the current translators string locale
func (l *khq_ML) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *khq_ML) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
