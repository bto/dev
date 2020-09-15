import * as ActionTypes from "./actionTypes";

let nextTodoId = 0;

export function addTodo(content: string) {
    return {
        type: ActionTypes.ADD_TODO,
        payload: {
            id: ++nextTodoId,
            content
        }
    };
}
