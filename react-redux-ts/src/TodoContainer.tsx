import * as React from "react";
import * as ReactRedux from "react-redux";
import * as Redux from "redux";

import { todoActionCreator } from "./action";
import { RootState } from "./store";
import TodoComponent from "./TodoComponent";

interface StateToProps {
    todos: string[];
}

interface DispatchToProps {
    addTodo: (todo: string) => void;
}

type Props = StateToProps & DispatchToProps;

class TodoContainer extends React.Component<Props, {}> {
    public render(): JSX.Element {
        const { todos } = this.props;
        return (
            <TodoComponent todos={todos} onClickAddButton={this.onClickAddButton} />
        )
    }

    private onClickAddButton = (todo: string): void => {
        const { addTodo } = this.props;
        addTodo(todo);
    }
}

const mapStateToProps = (state: RootState): StateToProps => {
    const { todoState } = state;
    return {
        todos: todoState.todos,
    };
}

const mapDispatchToProps = (dispatch: Redux.Dispatch<Redux.Action>): DispatchToProps => {
    return {
        addTodo: (todo: string) => {
            dispatch(todoActionCreator.addTodoAction(todo));
        },
    };
}

export default ReactRedux.connect<StateToProps, DispatchToProps, RootState>(
    mapStateToProps,
    mapDispatchToProps
)(TodoContainer);
