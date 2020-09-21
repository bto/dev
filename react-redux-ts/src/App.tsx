import * as React from 'react';
import * as ReactRedux from "react-redux";

import store from "./store";
import TodoContainer from "./TodoContainer";
import './App.css';

function App() {
    return (
        <ReactRedux.Provider store={store}>
            <TodoContainer />
        </ReactRedux.Provider>
    );
}

export default App;
