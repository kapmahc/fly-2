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

const history = createHistory()

const middleware = routerMiddleware(history)

const store = createStore(
  combineReducers({
    ...reducers,
    router: routerReducer
  }),
  applyMiddleware(middleware)
)

function main(user) {
  addLocaleData(user.data)
  ReactDOM.render(
    <Provider store={store}>
      <ConnectedRouter history={history}>
        <LocaleProvider locale={user.antd}>
          <IntlProvider locale={user.locale} messages={user.messages}>
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
