package bo_CN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bo_CN struct {
	locale string
}

// New returns a new instance of translator for the 'bo_CN' locale
func New() locales.Translator {
	return &bo_CN{
		locale: "bo_CN",
	}
}

// Locale returns the current translators string locale
func (l *bo_CN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bo_CN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
