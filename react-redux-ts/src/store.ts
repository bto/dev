import * as Redux from "redux";

import reducer from "./reducer";

export interface TodoState {
    todos: string[];
}

export interface RootState {
    todoState: TodoState;
}

declare module "react-redux" {
    interface DefaultRootState extends RootState {}
}

const store: Redux.Store<TodoState, Redux.Action> = Redux.createStore(reducer);

export default store;
