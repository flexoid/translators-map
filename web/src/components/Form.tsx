import { useState, useEffect } from "react"
import { Flex, Box, Select, Link, Spinner } from "@chakra-ui/react"
import { ExternalLinkIcon } from "@chakra-ui/icons"
import { Language, Translator } from "../lib/api"

type FormProps = {
  currentLanguage: string | null
  languages: Language[]
  visibleTranslators: Translator[]
  loading: boolean
  onLangChange: (lang: string) => void
}

function Form({
  languages,
  visibleTranslators,
  loading,
  onLangChange,
}: FormProps) {
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
          placeholder="Select language"
          onChange={(e) => onLangChange(e.target.value)}
        >
          {languages.map((language, index) => {
            return <option key={index}>{language.language}</option>
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
