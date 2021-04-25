import React from 'react'
import {Route, Switch} from 'react-router-dom'
import LandingPageComponent from '../views/landingPage/LandingPage.js'
import {connect} from 'react-redux'

import {
    ROUTE_HOME,
} from '../constants/routes'

class AppContainer extends React.Component {
   render() {
    return (
      <div style={{
        height: '100%',
      }}>
        <Switch>
          <Route path={ROUTE_HOME} component={LandingPageComponent}/>
        </Switch>
      </div>
    )
  }
}

const mapDispatchToProps = dispatch => {
  return {}
};

const mapStateToProps = state => {
  return {}
};
export default connect(mapStateToProps, mapDispatchToProps)(AppContainer)

