package ii_CN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ii_CN struct {
	locale string
}

// New returns a new instance of translator for the 'ii_CN' locale
func New() locales.Translator {
	return &ii_CN{
		locale: "ii_CN",
	}
}

// Locale returns the current translators string locale
func (l *ii_CN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ii_CN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
