import React from "react"
import ReactDOM from "react-dom/client"
import { i18n } from "@lingui/core"
import { I18nProvider } from "@lingui/react"
import { detect, fromNavigator, fromPath } from "@lingui/detect-locale"
import { messages as enMessages } from "./locales/en/messages"
import { messages as plMessages } from "./locales/pl/messages"
import { messages as ruMessages } from "./locales/ru/messages"
import App from "./App"
import "./index.css"

i18n.load({
  en: enMessages,
  pl: plMessages,
  ru: ruMessages,
})

let detectedLocale = detect(fromPath(0), fromNavigator())
if (detectedLocale) {
  detectedLocale = detectedLocale.split("-")[0]
} else {
  detectedLocale = "en"
}
i18n.activate(detectedLocale)

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <I18nProvider i18n={i18n}>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </I18nProvider>
)
