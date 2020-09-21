import * as Redux from "redux";
import { AddTodoAction, TodoAction, TodoActionType } from "./action";
import { RootState, TodoState } from "./store";

const initTodoState: TodoState = {
    todos: [],
};

const todoReducer: Redux.Reducer<TodoState, TodoAction> = (
    state: TodoState = initTodoState,
    action: TodoAction,
): TodoState => {
    switch (action.type) {
        case TodoActionType.ADD_TODO:
            const addTodoAction: AddTodoAction = action;
            return {
                ...state,
                todos: state.todos.concat([addTodoAction.payload.todo]),
            };
        default:
            return state;
    }
}

const reducer: Redux.Reducer<RootState, TodoAction> = Redux.combineReducers({
    todoState: todoReducer,
});

export default reducer;
