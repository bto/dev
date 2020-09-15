import React from "react";
import AddTodo from "./AddTodo";
import "./styles.css";

export default function App() {
    return (
        <div className="todo-app">
            <h1>Todo List</h1>
            <AddTodo />
        </div>
    )
}