import * as React from "react";
import * as ReactRedux from "react-redux";
import { addTodo } from "../redux/actions";

class Component extends React.Component {
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
