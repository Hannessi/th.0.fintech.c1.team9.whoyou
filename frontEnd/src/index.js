import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import {createStore, applyMiddleware} from 'redux'
import thunk from 'redux-thunk'
import config from 'react-global-configuration'
import Root from './root/Root'
import * as serviceWorker from './serviceWorker'
import reducers from './reducers/index'

const initialState = {};

const createStoreWithMiddleware = applyMiddleware(
  thunk,
)(createStore);

let store = createStoreWithMiddleware(
  reducers,
  initialState,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__());

let cfg = {};

// cfg.golang = 'http://0.0.0.0:9000/api'

// cfg.golang = 'http://127.0.0.1:9002/api'
cfg.golang = 'http://13.244.141.27:9002/api';

config.set(cfg);

ReactDOM.render(<Root store={store}/>, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();

