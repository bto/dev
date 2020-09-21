import React from 'react';
import TodoComponent from "./TodoComponent";
import logo from './logo.svg';
import './App.css';

function App() {
    return (
        <div>
            <TodoComponent
                onClickAddButton={(todo: string): void => {
                    console.log(todo);
                }}
            />
        </div>
    );
}

export default App;
