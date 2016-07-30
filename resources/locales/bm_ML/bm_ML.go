package bm_ML

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bm_ML struct {
	locale string
}

// New returns a new instance of translator for the 'bm_ML' locale
func New() locales.Translator {
	return &bm_ML{
		locale: "bm_ML",
	}
}

// Locale returns the current translators string locale
func (l *bm_ML) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bm_ML) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
