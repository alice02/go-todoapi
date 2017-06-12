import React from 'react';
import { render } from 'react-dom';
import { Provider } from "react-redux";
import injectTapEventPlugin from "react-tap-event-plugin";
import App from "./containers/App";
import configureStore from "./store/configureStore";
import { getTodosIfNeeded } from "./actions";

const store = configureStore()
store.dispatch(getTodosIfNeeded())
injectTapEventPlugin();

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.querySelector('#root')
);
