import { connect } from "react-redux";
import TodoListComponent from "../components/TodoList";


function mapStateToProps(state) {
  return {
    todos: state.todoList.todos
  };
}

export default connect(mapStateToProps)(TodoListComponent);
