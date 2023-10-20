import { useState, useEffect } from "react"
import { Language, Translator } from "../lib/api"
import { t } from "@lingui/macro"
import { useLingui } from "@lingui/react"
import languageNamesList from "../locales/language-names.json"
import Select from "@mui/joy/Select"
import Option from "@mui/joy/Option"
import { Box } from "@mui/joy"

type FormProps = {
  currentLanguage: string | null
  languages: Language[]
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

function Form({ languages, onLangChange }: FormProps) {
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
    <Box>
      <Select
        placeholder={t`Select language`}
        onChange={(_, value: string | null) => value && onLangChange(value)}
      >
        {languageItems.map((item, index) => {
          return (
            <Option key={index} value={item.origName}>
              {item.prettyName}
            </Option>
          )
        })}
      </Select>
    </Box>
  )
}

export default Form
