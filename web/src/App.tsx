import { useState, useEffect } from "react"
import { Flex, Box, Heading, Spacer, VStack } from "@chakra-ui/react"
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
      minWidth="max-content"
      h={{ base: "auto", md: "100vh" }}
      direction={{ base: "column", md: "row" }}
    >
      <VStack p={2} minWidth="max-content" padding={4}>
        <Heading size="md">Translators</Heading>

        <Form
          currentLanguage={currentLanguage}
          languages={languages}
          onLangChange={handleLangChange}
        />
      </VStack>
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
