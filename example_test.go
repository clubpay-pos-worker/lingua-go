/*
 * Copyright © 2021-present Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lingua_test

import (
	"fmt"

	"github.com/clubpay-pos-worker/lingua-go"
)

func Example_basic() {
	languages := []lingua.Language{
		lingua.English,
		lingua.French,
		lingua.German,
		lingua.Spanish,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	if language, exists := detector.DetectLanguageOf("languages are awesome"); exists {
		fmt.Println(language)
	}

	// Output: English
}

// By default, Lingua uses lazy-loading to load only those language models on
// demand which are considered relevant by the rule-based filter engine. For web
// services, for instance, it is rather beneficial to preload all language models
// into memory to avoid unexpected latency while waiting for the service response.
// If you want to enable the eager-loading mode, you can do it as seen below.
// Multiple instances of LanguageDetector share the same language models in
// memory which are accessed asynchronously by the instances.
func Example_eagerLoading() {
	lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()
}

// There might be classification tasks where you know beforehand that your language
// data is definitely not written in Latin, for instance. The detection accuracy
// can become better in such cases if you exclude certain languages from the
// decision process or just explicitly include relevant languages.
func Example_builderApi() {
	// Include all languages available in the library.
	lingua.NewLanguageDetectorBuilder().FromAllLanguages()

	// Include only languages that are not yet extinct (= currently excludes Latin).
	lingua.NewLanguageDetectorBuilder().FromAllSpokenLanguages()

	// Include only languages written with Cyrillic script.
	lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithCyrillicScript()

	// Exclude only the Spanish language from the decision algorithm.
	lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithout(lingua.Spanish)

	// Only decide between English and German.
	lingua.NewLanguageDetectorBuilder().FromLanguages(lingua.English, lingua.German)

	// Select languages by ISO 639-1 code.
	lingua.NewLanguageDetectorBuilder().FromIsoCodes639_1(lingua.EN, lingua.DE)

	// Select languages by ISO 639-3 code.
	lingua.NewLanguageDetectorBuilder().FromIsoCodes639_3(lingua.ENG, lingua.DEU)
}
