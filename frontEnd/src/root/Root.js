import React from "react"
import PropTypes from "prop-types"
import { Provider } from "react-redux"
import AppContainer from './AppContainer'
import {BrowserRouter as Router, Route, Switch} from "react-router-dom"

const Root = ({ store }) => (
  <Provider store={store}>
    <Router>
      <div>
        <Switch>
          <Route path="/" component={AppContainer}/>
        </Switch>
      </div>
    </Router>
  </Provider>
);

Root.propTypes = {
  store: PropTypes.object.isRequired,
};

export default Root
