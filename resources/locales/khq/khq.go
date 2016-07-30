package khq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type khq struct {
	locale string
}

// New returns a new instance of translator for the 'khq' locale
func New() locales.Translator {
	return &khq{
		locale: "khq",
	}
}

// Locale returns the current translators string locale
func (l *khq) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *khq) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
