package sg_CF

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sg_CF struct {
	locale string
}

// New returns a new instance of translator for the 'sg_CF' locale
func New() locales.Translator {
	return &sg_CF{
		locale: "sg_CF",
	}
}

// Locale returns the current translators string locale
func (l *sg_CF) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sg_CF) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
