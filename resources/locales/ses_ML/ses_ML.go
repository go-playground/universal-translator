package ses_ML

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ses_ML struct {
	locale string
}

// New returns a new instance of translator for the 'ses_ML' locale
func New() locales.Translator {
	return &ses_ML{
		locale: "ses_ML",
	}
}

// Locale returns the current translators string locale
func (l *ses_ML) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ses_ML) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
