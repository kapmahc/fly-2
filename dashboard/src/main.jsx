import React from 'react'
import ReactDOM from 'react-dom'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import createHistory from 'history/createBrowserHistory'
import { ConnectedRouter, routerReducer, routerMiddleware } from 'react-router-redux'
import { LocaleProvider } from 'antd'
import { IntlProvider, addLocaleData } from 'react-intl'

import reducers from './reducers'
import plugins from './plugins'
import i18n from './i18n'

const userLocale = i18n()
addLocaleData(userLocale.data)

const history = createHistory()

const middleware = routerMiddleware(history)

const store = createStore(
  combineReducers({
    ...reducers,
    router: routerReducer
  }),
  applyMiddleware(middleware)
)

function main() {
  ReactDOM.render(
    <Provider store={store}>
      <ConnectedRouter history={history}>
        <LocaleProvider locale={userLocale.antd}>
          <IntlProvider locale={userLocale.locale} messages={userLocale.messages}>
            <div>
              {plugins.routes}
            </div>
          </IntlProvider>
        </LocaleProvider>
      </ConnectedRouter>
    </Provider>,
    document.getElementById('root')
  );
}

export default main
