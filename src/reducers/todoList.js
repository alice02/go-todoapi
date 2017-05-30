import {
  FETCH_TODOS,
  RECEIVE_TODOS,
  ADD_TODO,
  ADD_FAILED,
  SAVE_TODO
} from "../actions";

function todoList(state = {isFetching: false, message: "", todos: []}, action) {
  switch(action.type) {
    case FETCH_TODOS:
      return Object.assign({}, state, {
        isFetching: true
      });
    case RECEIVE_TODOS:
      return Object.assign({}, state, {
        isFetching: false,
        todos: action.todos
      });
    case ADD_TODO:
      return Object.assign({}, state, {
        message: "",
        todos: [
          ...state.todos,
          action.todo
        ]
      });
    case ADD_FAILED:
      console.log(action.message);
      return Object.assign({}, state, {
        message: action.message
      });
    case SAVE_TODO:
      return Object.assign({}, state, {
        todos: state.todos
      });
    default:
      return state;
  }
}

export default todoList;
