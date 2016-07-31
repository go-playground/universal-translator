package ses_ML

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ses_ML struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ses_ML' locale
func New() locales.Translator {
	return &ses_ML{
		locale:  "ses_ML",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ses_ML) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ses_ML'
func (t *ses_ML) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ses_ML'
func (t *ses_ML) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
