import * as React from "react";

type Props = {
    onClickAddButton: (todo: string) => void;
    todos: string[];
};

type State = {
    text: string;
};

export default class extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props);
        this.state = {
            text: "",
        };
    }

    public render() {
        const { todos } = this.props;
        const { text } = this.state;
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
        const { onClickAddButton } = this.props;
        const { text } = this.state;
        onClickAddButton(text);
    }
}
