import React from 'react';
import PropTypes from "prop-types";

class TodoList extends React.Component {
  render() {
    return (
      <ul>
        {this.props.todos.map((todo, index) =>
          <li key={index}>{todo.description}</li>
        )}
      </ul>
    );
  }
}

TodoList.propTypes = {
  todos: PropTypes.array,
};

export default TodoList;
