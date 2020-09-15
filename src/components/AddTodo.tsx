import * as React from "react";
import * as ReactRedux from "react-redux";
import { addTodo } from "../redux/actions";

type Props = {};

type State = {
    input: string;
};

class Component extends React.Component<Props, State> {
    state = {
        input: "",
    };

    updateInput(input: string) {
        this.setState({ input });
    }

    handleAddTodo() {
        this.props.addTodo(this.state.input);
        this.setState({ input: "" });
    }

    render() {
        return (
            <div>
                <input
                    onChange={e => this.updateInput(e.target.value)}
                    value={this.state.input}
                />
                <button className="add-todo" onClick={this.handleAddTodo}>
                    Add Todo
                </button>
            </div>
        )
    };
}

export default ReactRedux.connect(
    null,
    { addTodo }
)(Component);
