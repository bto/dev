import * as Redux from "redux";

export enum TodoActionType {
    ADD_TODO = "ADD_TODO",
}

export interface AddTodoAction extends Redux.Action {
    type: TodoActionType.ADD_TODO;
    payload: {
        todo: string;
    };
}

export type TodoAction = AddTodoAction

export interface ITodoActionCreator {
    addTodoAction(todo: string): AddTodoAction;
}

class TodoActionCreator implements ITodoActionCreator {
    public addTodoAction = (todo: string): AddTodoAction => {
        return {
            type: TodoActionType.ADD_TODO,
            payload: {
                todo
            },
        };
    };
}

export const todoActionCreator: ITodoActionCreator = new TodoActionCreator();
