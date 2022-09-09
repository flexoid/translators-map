package scraper

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
	"testing"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestProcessTable(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fixtureFileContent(t, "list1.html"))
	}))
	defer ts.Close()

	c := colly.NewCollector()
	logger := zap.NewNop()
	language := Language{Name: "Polski", Code: 1}

	var translators []Translator

	c.OnHTML("table.tabelkaszara", func(e *colly.HTMLElement) {
		processTable(logger.Sugar(), e, language, func(translator Translator) {
			translators = append(translators, translator)
		})
	})

	c.Visit(ts.URL)
	c.Wait()

	assert.Equal(t, 14, len(translators))

	assert.Equal(t, "Tymoteusz Sulewski dr hab. nauk hum", translators[0].Name)
	assert.Equal(t, "Poprzeczna 42A 30-006 Kraków", translators[0].Address)
	assert.Equal(t, "", translators[0].Contacts)
	assert.Equal(t, "https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/translator,912398891767.html", translators[0].DetailsURL)

	assert.Equal(t, "Wiktoria Walczak", translators[1].Name)
	assert.Equal(t, "Wąska 73 37-310 Nowa Sarzyna, skr. poczt. 25", translators[1].Address)
	assert.Equal(t, "Tel:123-456-789 Email: wiktoria.walczak@gmail.com", translators[1].Contacts)
	assert.Equal(t, "https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/translator,742938749238.html", translators[1].DetailsURL)

	assert.Equal(t, "Ignacy Wieczorek", translators[2].Name)
	assert.Equal(t, "Wschodnia 20 82-500 Kwidzyn", translators[2].Address)
	assert.Equal(t, "Tel:(055) 123 45 67, 0601 001 123", translators[2].Contacts)
	assert.Equal(t, "https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/translator,4978238467.html", translators[2].DetailsURL)

	assert.Equal(t, "Joanna Kowalska", translators[8].Name)
	assert.Equal(t, "Avenue des Rogations 31, 1200 Bruksela, Belgia", translators[8].Address)
	assert.Equal(t, "", translators[8].Contacts)
	assert.Equal(t, "https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/translator,789883493.html", translators[8].DetailsURL)

	assert.Equal(t, "Justyna Wróbel", translators[13].Name)
	assert.Equal(t, "Telimeny 91/4 30-838 Kraków", translators[13].Address)
	assert.Equal(t, "Tel:999-351-818 Email: justyna@example.com", translators[13].Contacts)
	assert.Equal(t, "https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/translator,918374671.html", translators[13].DetailsURL)

	// TODO: Add assertions for the rest of the translators.
}

func fixtureFileContent(t *testing.T, filename string) []byte {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get current file for fixture loading")
	}

	fixturePath := path.Join(path.Dir(currentFile), "../../testdata/"+filename)
	fixturePage, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		t.Fatal(err)
	}

	return fixturePage
}
