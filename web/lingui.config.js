/** @type {import('@lingui/conf').LinguiConfig} */
module.exports = {
  locales: ["pl", "en", "ru"],
  catalogs: [
    {
      path: "src/locales/{locale}/messages",
      include: ["src"],
    },
  ],
  format: "po",
  sourceLocale: "en",
  fallbackLocales: {
    "ru-RU": "ru",
    "pl-PL": "pl",
    default: "en",
  },
  extractBabelOptions: {
    presets: ["@babel/preset-typescript"],
  },
}
