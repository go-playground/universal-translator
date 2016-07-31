package nl_SR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_SR struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nl_SR' locale
func New() locales.Translator {
	return &nl_SR{
		locale:  "nl_SR",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nl_SR) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nl_SR'
func (t *nl_SR) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nl_SR'
func (t *nl_SR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
