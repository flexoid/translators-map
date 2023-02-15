import { useState, useEffect } from "react"
import { Flex, Box, Heading, Text, Link } from "@chakra-ui/react"
import { ExternalLinkIcon } from "@chakra-ui/icons"
import { Wrapper, Status } from "@googlemaps/react-wrapper"
import ReactGA from "react-ga4"
import { Config } from "./lib/api"

import Map from "./components/Map"
import "./App.css"
import Form from "./components/Form"
import { Language, Translator } from "./lib/api"
import { Trans, t } from "@lingui/macro"

function App() {
  const [config, setConfig] = useState<Config | null>(null)
  const [currentLanguage, setCurrentLanguage] = useState<string | null>(null)
  const [languages, setLanguages] = useState<Language[]>([])
  const [translators, setTranslators] = useState<Translator[]>([])
  const [visibleTranslators, setVisibleTranslators] = useState<Translator[]>([])
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    fetch("/api/config")
      .then((response) => {
        return response.json()
      })
      .then((config) => {
        setConfig(config)
      })
  }, [])

  useEffect(() => {
    fetch("/api/languages")
      .then((response) => {
        return response.json()
      })
      .then((languages) => {
        setLanguages(languages)
      })
  }, [])

  useEffect(() => {
    if (config && config.google_analytics_id) {
      ReactGA.initialize(config.google_analytics_id)
    }
  }, [config])

  useEffect(() => {
    document.title = t({ id: "title" })
  }, [])

  const render = (status: Status) => {
    return <h1>{status}</h1>
  }

  const handleLangChange = (lang: string) => {
    setCurrentLanguage(lang)
  }

  useEffect(() => {
    setTranslators([])

    if (!currentLanguage) {
      return
    }

    setLoading(true)
    fetch(`/api/translators?lang=${currentLanguage}`)
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        setTranslators(data)
        setLoading(false)
      })
  }, [currentLanguage])

  return (
    <Flex height={{ base: "auto", md: "100vh" }} direction="column">
      <Flex direction={{ base: "column", md: "row" }} flex="auto">
        <Flex
          direction="column"
          p="4"
          alignItems="center"
          width={{ base: "auto", md: "md" }}
          flex="none"
        >
          <Heading size="md" pt={4} flex="none" textAlign="center">
            <Trans id="title">Sworn translations on map</Trans>
          </Heading>
          <Text p="4" flex="none">
            <Trans id="subtitle">Find a certified Polish translator.</Trans>
          </Text>
          <Form
            currentLanguage={currentLanguage}
            languages={languages}
            visibleTranslators={visibleTranslators}
            loading={loading}
            onLangChange={handleLangChange}
          />
        </Flex>

        <Box w="full" h={{ base: "xl", md: "full" }} flex="auto">
          {config && (
            <Wrapper
              apiKey={config.maps_js_api_key}
              render={render}
              libraries={["geometry"]}
            >
              <Map
                center={{ lat: 51.919438, lng: 19.145136 }}
                zoom={6}
                style={{ height: "100%" }}
                translators={translators}
                onVisibleTranslatorsChange={setVisibleTranslators}
              />
            </Wrapper>
          )}
        </Box>
      </Flex>

      <Box flex="none" p={4}>
        <Text fontSize="sm" align="center" margin="auto">
          <Trans id="disclamer">
            All data used on this site is taken from the{" "}
            <Link
              color="teal.500"
              href="https://arch-bip.ms.gov.pl/pl/rejestry-i-ewidencje/tlumacze-przysiegli/lista-tlumaczy-przysieglych/search.html"
              isExternal
            >
              <span style={{ whiteSpace: "nowrap" }}>
                Bulletin of Public information archive{" "}
                <ExternalLinkIcon mx="2px" />
              </span>
            </Link>{" "}
            of the Ministry of Justice of the Republic of Poland.
            <br />
            The data is provided "as is" without warranty of any kind for
            informational purposes only.
          </Trans>
        </Text>
      </Box>
    </Flex>
  )
}

export default App
