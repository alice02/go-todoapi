import React from "react";
import { connect } from "react-redux";
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

import { createTodo } from "../actions";
import Header from "../components/Header";
import TodoInput from "../components/TodoInput";
import TodoList from "../components/TodoList";

class App extends React.Component {
  render() {
    return (
      <MuiThemeProvider>
        <div>
          <Header />
          <TodoInput
            isFetching={this.props.isFetching}
            message={this.props.message}
            save={this.props.save}
          />
          <TodoList todos={this.props.todos} />
        </div>
      </MuiThemeProvider>
    );
  }
}

function mapStateToProps(state) {
  return {
    isFetching: state.todoList.isFetching,
    message: state.todoList.message,
    todos: state.todoList.todos
  };
}

function mapDispatchToProps(dispatch) {
  return {
    save: (todo) => { dispatch(createTodo(todo)) },
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(App);
