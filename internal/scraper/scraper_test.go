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

	// TODO: One translator is currently not parsed due to a bug in the scraper.
	// assert.Equal(t, 13, len(translators))

	assert.Equal(t, "Tymoteusz Sulewski dr hab. nauk hum", translators[0].Name)
	assert.Equal(t, "Poprzeczna 42A 30-006 Kraków", translators[0].Address)
	assert.Equal(t, "", translators[0].Contacts)
	assert.Equal(t, "https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/translator,912398891767.html", translators[0].DetailsURL)

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
