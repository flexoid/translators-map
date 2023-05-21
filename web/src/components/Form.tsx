import { useState, useEffect } from "react"
import { Flex, Box, Select, Link, Spinner } from "@chakra-ui/react"
import { ExternalLinkIcon } from "@chakra-ui/icons"
import { Language, Translator } from "../lib/api"
import { Trans, t } from "@lingui/macro"
import { useLingui } from "@lingui/react"
import languageNamesList from "../locales/language-names.json"

type FormProps = {
  currentLanguage: string | null
  languages: Language[]
  visibleTranslators: Translator[]
  loading: boolean
  onLangChange: (lang: string) => void
}

interface LanguageMap {
  [polishName: string]: {
    [languageCode: string]: string
  }
}

interface LanguageItem {
  origName: string
  prettyName: string
}

const languageMap: LanguageMap = languageNamesList

function Form({
  languages,
  visibleTranslators,
  loading,
  onLangChange,
}: FormProps) {
  const { i18n } = useLingui()
  const [languageItems, setlanguageItems] = useState<LanguageItem[]>([])

  useEffect(() => {
    const languageItems = languages.map((language) => {
      // Get translated language name from the file.
      let prettyName =
        languageMap[language.language]?.[i18n.locale] || language.language
      // Capitalize name.
      prettyName = prettyName.charAt(0).toUpperCase() + prettyName.slice(1)

      return {
        origName: language.language,
        prettyName: prettyName,
      }
    })

    // Sort languages by pretty name.
    languageItems.sort((a, b) => {
      if (a.prettyName < b.prettyName) {
        return -1
      }
      if (a.prettyName > b.prettyName) {
        return 1
      }
      return 0
    })

    setlanguageItems(languageItems)
  }, [languages, i18n.locale])

  return (
    <Flex
      direction="column"
      width="full"
      justifyContent="center"
      flex="1 1 0"
      minHeight={{ base: "auto", md: 0 }}
      maxHeight={{ base: "80", md: "fit-content" }}
    >
      <Box p={2} flex="none">
        <Select
          placeholder={t`Select language`}
          onChange={(e) => onLangChange(e.target.value)}
        >
          {languageItems.map((item, index) => {
            return (
              <option key={index} value={item.origName}>
                {item.prettyName}
              </option>
            )
          })}
        </Select>
      </Box>

      {loading && <Spinner size="lg" flex="none" alignSelf="center" />}

      <Flex direction="column" p={2} flex="auto" overflowY="auto">
        {visibleTranslators.map((translator, index) => {
          return (
            <Box key={index} p={3} flex="none">
              <Link href={translator.details_url} isExternal>
                {index + 1}. {translator.address} <ExternalLinkIcon mx="2px" />
              </Link>
            </Box>
          )
        })}
      </Flex>
    </Flex>
  )
}

export default Form
