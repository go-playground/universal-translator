package ut

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/locales-list"
)

// UniversalTranslator holds all locale & translation data
type UniversalTranslator struct {
	translators map[string]Translator
	fallback    Translator
}

// New returns a new UniversalTranslator instance set with
// the fallback locale and locales it should support
func New(fallback string, supportedLocales ...string) (*UniversalTranslator, error) {

	t := &UniversalTranslator{
		translators: make(map[string]Translator),
	}

	lcMapping := make(localeslist.LocaleMap)

	for loc, tf := range localeslist.Map() {
		lcMapping[strings.ToLower(loc)] = tf
	}

	var ok bool
	var lc string
	var tf localeslist.LocaleFunc
	fallbackLower := strings.ToLower(fallback)

	for _, loc := range supportedLocales {

		lc = strings.ToLower(loc)

		tf, ok = lcMapping[lc]
		if !ok {
			return nil, fmt.Errorf("locale '%s' could not be found", loc)
		}

		if fallbackLower == lc {
			t.fallback = newTranslator(tf())
			t.translators[lc] = t.fallback
			continue
		}

		t.translators[lc] = newTranslator(tf())
	}

	if t.fallback == nil {
		return nil, fmt.Errorf("fallback locale '%s' could not be found", fallback)
	}

	return t, nil
}

// FindTranslator trys to find a Translator based on an array of locales
// and returns the first one it can find, otherwise returns the
// fallback translator.
func (t *UniversalTranslator) FindTranslator(locales ...string) (trans Translator) {

	var ok bool

	for _, locale := range locales {

		if trans, ok = t.translators[strings.ToLower(locale)]; ok {
			return
		}
	}

	return t.fallback
}

// GetTranslator returns the specified translator for the given locale,
// or fallback if not found
func (t *UniversalTranslator) GetTranslator(locale string) Translator {

	if t, ok := t.translators[strings.ToLower(locale)]; ok {
		return t
	}

	return t.fallback
}
