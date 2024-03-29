package scraper

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

const baseUrl = "https://arch-bip.ms.gov.pl"

var idInUrlRegex = regexp.MustCompile(`\/translator\,(\d+)\.html`)

type Translator struct {
	ID         int
	Name       string
	Address    string
	Contacts   string
	DetailsURL string
	Language
}

type Language struct {
	Name string
	Code int
}

func ScrapeTranslators(logger *zap.SugaredLogger, language Language, cb func(Translator)) (outerErr error) {
	c := colly.NewCollector()
	c.SetRequestTimeout(60 * time.Second)
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	logger = logger.With("language", language.Name)

	var nextPage string

	c.OnHTML("html", func(e *colly.HTMLElement) {
		nextPage = nextPageLink(e)
	})

	c.OnHTML("table.tabelkaszara", func(e *colly.HTMLElement) {
		processTable(logger, e, language, cb)
	})

	c.OnRequest(func(r *colly.Request) {
		logger.Debugf("Visiting %s", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		outerErr = err
	})

	url := fmt.Sprintf(
		"%s/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/search.html?Language=%d",
		baseUrl, language.Code)

	outerErr = c.Visit(url)
	if outerErr != nil {
		return
	}

	c.Wait()

	for nextPage != "" {
		url = baseUrl + nextPage
		nextPage = ""

		outerErr = c.Visit(url)
		if outerErr != nil {
			return
		}
	}

	return
}

func ScrapeLanguages(logger *zap.SugaredLogger) (languages []Language, outerErr error) {
	c := colly.NewCollector()
	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})

	c.OnRequest(func(r *colly.Request) {
		logger.Debugf("Visiting %s", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		outerErr = err
	})

	c.OnHTML(".formularz select[name=Language] option", func(e *colly.HTMLElement) {
		lang := Language{
			Name: e.Text,
		}

		value := e.Attr("value")
		if value == "" {
			return
		}

		code, err := strconv.Atoi(value)
		if err != nil {
			logger.Errorf("Failed to parse language code: %v", err)
			return
		}
		lang.Code = code

		languages = append(languages, lang)
	})

	outerErr = c.Visit(baseUrl + "/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/search.html")
	if outerErr != nil {
		return
	}

	c.Wait()

	return
}

func processTable(logger *zap.SugaredLogger, e *colly.HTMLElement, language Language, cb func(Translator)) {
	logger = logger.With("page", e.Request.URL)

	var rows []*colly.HTMLElement
	e.ForEach("tr", func(i int, h *colly.HTMLElement) {
		rows = append(rows, h)
	})
	if len(rows) < 2 {
		return
	}

	for i, element := range rows[1:] {
		logger.Debugf("Parsing translator from table row %d", i+1)
		translator := Translator{Language: language}

		seqNum, err := extractSeqNumber(element)
		if err != nil {
			logger.Errorf("Failed to extract sequential number: %v", err)
			continue
		}

		// Having sequential number in logs is useful for debugging.
		logger := logger.With("seq_num", seqNum)

		link, err := extractLink(element)
		if err != nil {
			logger.Errorf("Failed to extract link: %v", err)
			continue
		}
		translator.DetailsURL = link

		translator.ID, err = extractIDFromURL(link)
		if err != nil {
			logger.Errorf("Failed to extract ID: %v", err)
			continue
		}

		logger = logger.With("translator_id", translator.ID)

		err = parseAddressCell(element, &translator)
		if err != nil {
			logger.Errorf("Failed to parse address cell: %v", err)
			continue
		}

		name, err := extractName(element)
		if err != nil {
			logger.Errorf("Failed to extract name: %v", err)
			continue
		}
		translator.Name = name

		cb(translator)
	}
}

func extractSeqNumber(tr *colly.HTMLElement) (string, error) {
	text := tr.ChildText("td:nth-child(1)")
	if text == "" {
		return "", errors.New("could not find number")
	}

	return text, nil
}

func extractName(tr *colly.HTMLElement) (string, error) {
	text := tr.ChildText("td:nth-child(2)")
	if text == "" {
		return "", errors.New("could not find name cell")
	}

	name := strings.Join(strings.Fields(text), " ")
	return name, nil
}

func extractLink(tr *colly.HTMLElement) (string, error) {
	link := tr.ChildAttr("td:nth-child(2) a", "href")
	if link == "" {
		return "", errors.New("could not find link element")
	}

	return baseUrl + link, nil
}

func parseAddressCell(tr *colly.HTMLElement, t *Translator) error {
	text := tr.ChildText("td:nth-child(7)")
	if text == "" {
		return errors.New("could not find address cell")
	}

	addrLines := []string{}
	contactLines := []string{}
	stillAddr := true

	for _, rawLine := range strings.Split(text, "\n") {
		line := strings.Join(strings.Fields(rawLine), " ")
		if line == "" {
			continue
		}

		if strings.HasPrefix(strings.ToLower(line), "tel:") ||
			strings.HasPrefix(strings.ToLower(line), "fax:") ||
			strings.HasPrefix(strings.ToLower(line), "email:") {
			stillAddr = false
		}

		if stillAddr {
			addrLines = append(addrLines, line)
		} else {
			contactLines = append(contactLines, line)
		}
	}

	t.Address = strings.Join(addrLines, " ")
	if t.Address == "" {
		return errors.New("address is empty")
	}

	t.Contacts = strings.Join(contactLines, " ")

	return nil
}

func nextPageLink(e *colly.HTMLElement) string {
	return e.ChildAttr(".pager a.next", "href")
}

func extractIDFromURL(link string) (int, error) {
	matches := idInUrlRegex.FindStringSubmatch(link)
	if len(matches) != 2 {
		return 0, errors.New("could not extract ID from URL")
	}

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("ID is not a number: %s, %v", matches[1], err)
	}

	return id, nil
}
