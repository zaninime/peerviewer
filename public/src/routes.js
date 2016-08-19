import React from 'react';
import {Route, IndexRoute} from 'react-router';
import IndexPage from './views/indexPage/IndexPage';

export const makeRoutes = () => {
  return (
    <Route path='/'>
      {/* Lazy-loading */}
      <Route path="about" getComponent={(location, cb) => {
          require.ensure([], (require) => {
            const mod = require('./views/about/About');
            cb(null, mod.default);
          });
        }} />
      {/* inline loading */}
      <IndexRoute component={IndexPage} />
    </Route>
  );
};

export default makeRoutes;
