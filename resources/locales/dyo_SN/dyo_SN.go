package dyo_SN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dyo_SN struct {
	locale string
}

// New returns a new instance of translator for the 'dyo_SN' locale
func New() locales.Translator {
	return &dyo_SN{
		locale: "dyo_SN",
	}
}

// Locale returns the current translators string locale
func (l *dyo_SN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dyo_SN) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
