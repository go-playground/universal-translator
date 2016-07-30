package nl_SR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nl_SR struct {
	locale string
}

// New returns a new instance of translator for the 'nl_SR' locale
func New() locales.Translator {
	return &nl_SR{
		locale: "nl_SR",
	}
}

// Locale returns the current translators string locale
func (l *nl_SR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nl_SR) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
