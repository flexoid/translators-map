import { useState, useEffect } from "react"
import { Flex, Box, Select } from "@chakra-ui/react"
import { Language } from "../lib/api"

type FormProps = {
  currentLanguage: string | null
  languages: Language[]
  onLangChange: (lang: string) => void
}

function Form({ languages, onLangChange }: FormProps) {
  return (
    <Flex width="full" justifyContent="center">
      <Box p={2}>
        <Select
          placeholder="Select language"
          onChange={(e) => onLangChange(e.target.value)}
        >
          {languages.map((language, index) => {
            return <option key={index}>{language.language}</option>
          })}
        </Select>
      </Box>
    </Flex>
  )
}

export default Form
