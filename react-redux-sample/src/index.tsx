import React from "react";
import ReactDOM from "react-dom";
import * as ReactRedux from "react-redux";

import App from "./components/App";
import store from "./redux/store";

ReactDOM.render(
  <ReactRedux.Provider store={store}>
    <App />
  </ReactRedux.Provider>,
  document.getElementById("root")
);
