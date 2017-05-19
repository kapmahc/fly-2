import {LOCALE} from './constants'
import {enUS, zhHans} from './locales'

export default () => {
  switch (localStorage.getItem(LOCALE)) {
    case 'zh-Hans':
      return zhHans
    default:
    return enUS
  }
}
