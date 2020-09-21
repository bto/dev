import * as React from "react";

type Props = {
    onClickAddButton: (todo: string) => void;
};

type State = {
    text: string;
    todos: string[];
};

export default class extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);
        this.state = {
            text: "",
            todos: [],
        };
    }

    public render() {
        const { text, todos } = this.state;
        return (
            <div style={{ width: "500px", margin: "0 auth"}}>
                <h1>Todo List</h1>
                <input type="text" value={text} onChange={this.onTextChange} />
                <button onClick={this.onClickAddButton}>Add Todo</button>
                <ul>
                    {todos.map((todo, i) => (
                        <li key={i}>{todo}</li>
                    ))}
                </ul>
            </div>
        );
    }

    private onTextChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        this.setState({ text: e.currentTarget.value });
    };

    private onClickAddButton = () => {
        const { text, todos } = this.state;
        todos.push(text);
        this.setState({
            todos: todos.concat([text]),
        });
    }
}
