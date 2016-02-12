package ut

import (
	"errors"
	"sync"
)

var utrans *UniversalTranslator
var once sync.Once

type translators map[string]*Translator

// UniversalTranslator holds all locale Translator instances
type UniversalTranslator struct {
	translators map[string]*Translator
	fallback    *Translator
}

// newUniversalTranslator creates a new UniversalTranslator instance.
func newUniversalTranslator() *UniversalTranslator {
	return &UniversalTranslator{
		translators: make(translators),
	}
}

// SetFallback registers the fallback Translator instance when no matching Translator
// can be found via a given locale
func SetFallback(translator *Translator) {
	utrans.fallback = translator
}

// GetFallback return the universal translators fallback translation
func GetFallback() *Translator {
	return utrans.fallback
}

// FindTranslator trys to find a Translator based on an array of locales
// and returns the first one it can find, otherwise returns the
// fallback translator
func FindTranslator(locales []string) *Translator {

	for _, locale := range locales {

		if t, ok := utrans.translators[locale]; ok {
			return t
		}
	}

	return utrans.fallback
}

// GetTranslator returns the specified translator for the given locale,
// or error if not found
func GetTranslator(locale string) (*Translator, error) {

	if t, ok := utrans.translators[locale]; ok {
		return t, nil
	}

	return nil, errors.New("Translator with locale '" + locale + "' counld not be found.")
}

// RegisterLocale registers the a locale with ut
// initializes singleton + sets initial fallback language
func RegisterLocale(loc *Locale) {
	once.Do(func() {
		utrans = newUniversalTranslator()
	})

	if _, ok := utrans.translators[loc.Locale]; ok {
		return
	}

	t := newTranslator(loc)

	if utrans.fallback == nil {
		utrans.fallback = t
	}

	utrans.translators[loc.Locale] = t

	// have do once initialize singleton ut instance.
}
