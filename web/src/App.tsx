import "@fontsource/inter"

import { useState, useEffect } from "react"
import { Wrapper, Status } from "@googlemaps/react-wrapper"
import ReactGA from "react-ga4"
import { t } from "@lingui/macro"
import { useLingui } from "@lingui/react"
import { CssVarsProvider } from "@mui/joy/styles"
import CssBaseline from "@mui/joy/CssBaseline"
import Box from "@mui/joy/Box"
import Stack from "@mui/joy/Stack"

import { Config } from "./lib/api"
import Map from "./components/Map"
import "./App.css"
import Form from "./components/Form"
import { Language, Translator } from "./lib/api"
import Header from "./components/Header"
import HeaderSection from "./components/HeaderSection"
import Footer from "./components/Footer"
import Results from "./components/Results"

function App() {
  const [config, setConfig] = useState<Config | null>(null)
  const [currentLanguage, setCurrentLanguage] = useState<string | null>(null)
  const [languages, setLanguages] = useState<Language[]>([])
  const [translators, setTranslators] = useState<Translator[]>([])
  const [visibleTranslators, setVisibleTranslators] = useState<Translator[]>([])
  const [loading, setLoading] = useState<boolean>(false)
  const { i18n } = useLingui()

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
    document.title = t({
      id: "title",
      message: "Polish sworn translators on map",
    })
  }, [i18n.locale])

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
    <CssVarsProvider>
      <CssBaseline />
      <Header />

      <Box
        component="main"
        sx={{
          height: "calc(100vh - 55px)", // 55px is the height of the NavBar
          display: "grid",
          gridTemplateColumns: { xs: "auto", md: "40% 60%" },
          gridTemplateRows: { xs: "auto 1fr auto auto", md: "auto 1fr auto" },
        }}
      >
        <Stack
          sx={{
            backgroundColor: "background.surface",
            px: { xs: 2, md: 4 },
            py: 2,
            borderBottom: "1px solid",
            borderColor: "divider",
          }}
        >
          <HeaderSection />
          <Form
            currentLanguage={currentLanguage}
            languages={languages}
            onLangChange={handleLangChange}
          />
        </Stack>

        <Box
          sx={{
            gridRow: { md: "span 2" },
            width: "full",
            height: { xs: "40vh", md: "auto" },
            backgroundColor: "background.level1",
          }}
        >
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

        <Results visibleTranslators={visibleTranslators} loading={loading} />
        <Footer />
      </Box>
    </CssVarsProvider>
  )
}

export default App
