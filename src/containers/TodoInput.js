import { connect } from "react-redux";
import { createTodo }　from "../actions";
import TodoInputComponent from "../components/TodoInput";


function mapStateToProps(state) {
  return state;
}

function mapDispatchToProps(dispatch) {
  return {
    save: (todo) => { dispatch(createTodo(todo)) }
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(TodoInputComponent);
