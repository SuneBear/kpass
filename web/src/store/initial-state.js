import { cookie } from 'utils'
import { dictionaries, languages } from '../locales'

const getInitialLocale = () => {
  let initialLocale = cookie('lang')
  const defaultLocale = 'en-US'
  const languageCodes = languages.map(language => language.code)
  if (languageCodes.indexOf(initialLocale) < 0) {
    initialLocale = defaultLocale
  }
  cookie('lang', initialLocale)
  return initialLocale
}

/**
 * State Tree
 * @types: API, View (UI), Form, Context
 */
export const getInitialState = () => ({
  i18n: {
    locale: getInitialLocale(),
    translations: dictionaries
  }
})
