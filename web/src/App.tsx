import { useState, useEffect } from "react"
import { Flex, Box, Heading, Text, Spacer, Link } from "@chakra-ui/react"
import { ExternalLinkIcon } from "@chakra-ui/icons"
import { Wrapper, Status } from "@googlemaps/react-wrapper"

import Map from "./components/Map"
import "./App.css"
import Form from "./components/Form"
import { Language, Translator } from "./lib/api"

type Config = {
  maps_js_api_key: string
}

function App() {
  const [config, setConfig] = useState<Config | null>(null)
  const [currentLanguage, setCurrentLanguage] = useState<string | null>(null)
  const [languages, setLanguages] = useState<Language[]>([])
  const [translators, setTranslators] = useState<Translator[]>([])

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

  const render = (status: Status) => {
    return <h1>{status}</h1>
  }

  const handleLangChange = (lang: string) => {
    setCurrentLanguage(lang)
  }

  useEffect(() => {
    if (!currentLanguage) {
      setTranslators([])
      return
    }

    fetch(`/api/translators?lang=${currentLanguage}`)
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        setTranslators(data)
      })
  }, [currentLanguage])

  return (
    <Flex
      h={{ base: "auto", md: "100vh" }}
      direction={{ base: "column", md: "row" }}
    >
      <Flex
        direction="column"
        p="4"
        alignItems="center"
        // minWidth="align-content"
        maxWidth={{ base: "auto", md: "md" }}
      >
        <Heading size="md" pt={4}>
          Polish sworn translators map
        </Heading>
        <Text p="4" align="center">
          Find sworn translator from any language to Polish and vice versa.
        </Text>
        <Form
          currentLanguage={currentLanguage}
          languages={languages}
          onLangChange={handleLangChange}
        />
        <Spacer />

        <Text fontSize="sm" align="center">
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
          of the Ministry of Justice of the Republic of Poland. The data is
          provided "as is" without warranty of any kind for informational
          purposes only.
        </Text>
      </Flex>
      <Spacer />
      <Box w="full" h={{ base: "xl", md: "full" }}>
        {config && (
          <Wrapper apiKey={config.maps_js_api_key} render={render}>
            <Map
              center={{ lat: 52.237049, lng: 21.017532 }}
              zoom={8}
              style={{ height: "100%" }}
              translators={translators}
            />
          </Wrapper>
        )}
      </Box>
    </Flex>
  )
}

export default App
