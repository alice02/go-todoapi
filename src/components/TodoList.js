import React from 'react';

const TodoList = ({ todos }) => (
  <ul>
    {todos.map(todo =>
      <li>{todo.description}</li>
     )}
  </ul>
)

export default TodoList;
