import { dictionaries } from '../locales'

/**
 * State Tree
 * @types: API, View (UI), Form, Context
 */
export const initialState = {
  i18n: {
    locale: 'en-US',
    translations: dictionaries
  }
}
