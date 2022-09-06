export type Language = {
  language: string
}

type TranslatorLocation = {
  lat: number
  lng: number
}

export type Translator = {
  name: string
  address: string
  contacts: string
  location: TranslatorLocation
}
