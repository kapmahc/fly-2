import antdEn from 'antd/lib/locale-provider/en_US'
import dataEn from 'react-intl/locale-data/en'
import dataZh from 'react-intl/locale-data/zh'

import {LOCALE} from '../constants'

const enUS = {
  antd: antdEn,
  data: dataEn,
  locale: 'en-US'
}

const zhHans = {
  antd: null,
  data: dataZh,
  locale: 'zh-Hans'
}

export default () => {
  switch (localStorage.getItem(LOCALE)) {
    case 'zh-Hans':
      return zhHans
    default:
    return enUS
  }
}
