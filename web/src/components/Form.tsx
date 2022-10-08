import { useState, useEffect } from "react"
import { Flex, Box, Select, Link } from "@chakra-ui/react"
import { ExternalLinkIcon } from "@chakra-ui/icons"
import { Language, Translator } from "../lib/api"

type FormProps = {
  currentLanguage: string | null
  languages: Language[]
  visibleTranslators: Translator[]
  onLangChange: (lang: string) => void
}

function Form({ languages, visibleTranslators, onLangChange }: FormProps) {
  return (
    <Flex
      direction="column"
      width="full"
      justifyContent="center"
      flex="auto"
      overflowY="auto"
      maxHeight={{ base: "80", md: "none" }}
    >
      <Box p={2} flex="none">
        <Select
          placeholder="Select language"
          onChange={(e) => onLangChange(e.target.value)}
        >
          {languages.map((language, index) => {
            return <option key={index}>{language.language}</option>
          })}
        </Select>
      </Box>

      <Flex direction="column" p={2} flex="auto" overflowY="inherit">
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
