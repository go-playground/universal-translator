package vai_Vaii_LR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Vaii_LR struct {
	locale string
}

// New returns a new instance of translator for the 'vai_Vaii_LR' locale
func New() locales.Translator {
	return &vai_Vaii_LR{
		locale: "vai_Vaii_LR",
	}
}

// Locale returns the current translators string locale
func (l *vai_Vaii_LR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *vai_Vaii_LR) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
