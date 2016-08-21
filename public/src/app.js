/* globals __DEBUG__ */
import React from 'react';
import ReactDOM from 'react-dom';

import App from 'containers/App/App';

import {browserHistory} from 'react-router';
//import makeRoutes from './routes';

import * as api from './api/streams';
window['api'] = api;

const initialState = {};
import {configureStore} from './configureStore';
const {store, actions, history} = configureStore({initialState, historyType: browserHistory});

let render = (routerKey = null) => {
  const makeRoutes = require('./routes').default;
  const routes = makeRoutes(store);

  const mountNode = document.querySelector('#root');
  ReactDOM.render(
    <App history={history}
          store={store}
          actions={actions}
          routes={routes}
          routerKey={routerKey} />, mountNode);
};

if (__DEBUG__ && module.hot) {
  const renderApp = render;
  render = () => renderApp(Math.random());

  module.hot.accept('./routes', () => render());
}

render();
