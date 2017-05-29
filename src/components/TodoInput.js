import React from "react";
import ReactDOM from "react-dom";

class TodoInput extends React.Component {
  handleClick() {
    const description = ReactDOM.findDOMNode(this.refs.description).value;
    console.log(description);
    this.props.save(description);
  }

  render() {
    return (
      <div>
        <input ref="description" className="keyword" type="text" />
        <button onClick={this.handleClick.bind(this)}>Add</button>
      </div>
    );
  }
}

export default TodoInput;
