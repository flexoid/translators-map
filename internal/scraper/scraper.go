package scraper

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

const baseUrl = "https://arch-bip.ms.gov.pl"

var separatorOtherRegex = regexp.MustCompile("[\\pZ\\pC]")

type Translator struct {
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

	c.Visit(url)
	c.Wait()

	for nextPage != "" {
		url = baseUrl + nextPage
		nextPage = ""

		c.Visit(url)
	}

	return outerErr
}

func ScrapeLanguages(logger *zap.SugaredLogger) (languages []Language, outerErr error) {
	c := colly.NewCollector()

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

	c.Visit(baseUrl + "/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/search.html")
	c.Wait()

	return languages, nil
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

		address, err := extractAddress(element)
		if err != nil {
			logger.Errorf("Failed to extract address: %v", err)
			continue
		}
		translator.Address = address

		name, err := extractName(element)
		if err != nil {
			logger.Errorf("Failed to extract name: %v", err)
			continue
		}
		translator.Name = name

		link, err := extractLink(element)
		if err != nil {
			logger.Errorf("Failed to extract link: %v", err)
			continue
		}
		translator.DetailsURL = link

		contacts, err := extractContacts(element)
		if err != nil {
			logger.Errorf("Failed to extract contacts: %v", err)
			continue
		}
		translator.Contacts = contacts

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

func extractAddress(tr *colly.HTMLElement) (string, error) {
	text := tr.ChildText("td:nth-child(7)")
	if text == "" {
		return "", errors.New("could not find address cell")
	}

	parts := strings.Split(text, "\n")
	if len(parts) < 3 {
		return "", errors.New("not enough address parts")
	}

	// Extract important part of address cell.
	addr := separatorOtherRegex.ReplaceAllString(strings.Join(parts[:3], ""), " ")

	// Strip whitespaces.
	addr = strings.Join(strings.Fields(addr), " ")

	return addr, nil
}

func extractContacts(tr *colly.HTMLElement) (string, error) {
	text := tr.ChildText("td:nth-child(7)")
	if text == "" {
		return "", errors.New("could not extract contacts")
	}

	parts := strings.Split(text, "\n")
	if len(parts) < 4 {
		// No contact details.
		return "", nil
	}

	// Extract important part of address cell.
	contacts := separatorOtherRegex.ReplaceAllString(strings.Join(parts[3:], ""), " ")

	// Strip whitespaces.
	contacts = strings.Join(strings.Fields(contacts), " ")

	return contacts, nil
}

func nextPageLink(e *colly.HTMLElement) string {
	return e.ChildAttr(".pager a.next", "href")
}
