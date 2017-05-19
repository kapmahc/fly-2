import antdEn from 'antd/lib/locale-provider/en_US'
import dataEn from 'react-intl/locale-data/en'
import dataZh from 'react-intl/locale-data/zh'

import messagesEn from './en-US.json'
import messagesZhHans from './zh-Hans.json'

export const enUS = {
  messages: {...messagesEn},
  antd: antdEn,
  data: dataEn,
  locale: 'en-US'
}

export const zhHans = {
  messages: {...messagesZhHans},
  antd: null,
  data: dataZh,
  locale: 'zh-Hans'
}
