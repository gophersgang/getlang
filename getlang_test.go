package getlang

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEmptyStringFromReader(t *testing.T) {
	info := FromReader(strings.NewReader(""))
	assert.Equal(t, "und", info.LanguageCode())
}

func TestEnglishPhraseFromReader(t *testing.T) {
	info := FromReader(strings.NewReader("this is the language"))
	assert.Equal(t, "en", info.LanguageCode())
	assert.Equal(t, true, info.Confidence() > 0.90)
}

func TestEnglishPhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"We hold these truths to be self-evident, that all men are created equal",
		"en",
		0.95)
}

func TestGermanPhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Wir halten diese Wahrheiten für ausgemacht, daß alle Menschen gleich erschaffen worden",
		"de",
		0.95)
}

func TestGermanMixedEnglish(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Wenn wir jemand grüßen wollen, sagen wir 'How are you doing?'",
		"de",
		0.95)
}

func TestEnglishMixedGerman(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"If you wanted to greet someone in this language, you'd say 'wie geht es'",
		"en",
		0.65)
}

func TestEnglishMixedUkranian(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"the best thing to say is своїй гідності in my opinon.",
		"en",
		0.55)
}

func TestSpanishPhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Sostenemos como evidentes estas verdades: que los hombres son creados iguales",
		"es",
		0.75)
}

func TestPortuguesePhraseUSDI(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Consideramos estas verdades como autoevidentes, que todos os homens são criados iguais",
		"pt",
		0.95)
}

func TestPolishPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Wszyscy ludzie rodzą się wolni i równi w swojej godności i prawach",
		"pl",
		0.95)
}

func TestHungarianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Minden emberi lény szabadon születik és egyenlő méltósága és joga van",
		"hu",
		0.95)
}

func TestItalianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Tutti gli esseri umani nascono liberi ed eguali in dignità e diritti",
		"it",
		0.95)
}

func TestRussianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Все люди рождаются свободными и равными в своем достоинстве и правах",
		"ru",
		0.55)
}

func TestUkranianPhraseUDHR(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"Всі люди народжуються вільними і рівними у своїй гідності та правах",
		"uk",
		0.95)
}

func TestNonsense(t *testing.T) {
	ensureClassifiedWithConfidence(
		t,
		"wep lvna eeii vl jkk azc nmn iuah ppl zccl c%l aa1z",
		"und",
		0.95)
}

func ensureClassifiedWithConfidence(t *testing.T, text string, expectedLang string, minConfidence float64) {
	info := FromString(text)

	assert.Equal(t, expectedLang, info.LanguageCode(), "Misclassified text: "+text)
	assert.Equal(t, true, info.Confidence() > minConfidence)
}