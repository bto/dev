import React from "react";
import ReactDOM from "react-dom";
import * as ReactRedux from "react-redux";

import TodoApp from "./TodoApp";
import store from "./redux/store";

ReactDOM.render(
  <ReactRedux.Provider store={store}>
    <TodoApp />
  </ReactRedux.Provider>,
  document.getElementById("root")
);
