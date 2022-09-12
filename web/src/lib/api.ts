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
  details_url: string
  location: TranslatorLocation
}

export type Config = {
  maps_js_api_key: string
  google_analytics_id: string
}
