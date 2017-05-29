import {
  FETCH_TODOS,
  RECEIVE_TODOS,
  ADD_TODO,
  SAVE_TODO
} from "../actions";

function todoList(state = {isFetching: false, todos: []}, action) {
  switch(action.type) {
    case FETCH_TODOS:
      return state;
    case RECEIVE_TODOS:
      return state;
    case ADD_TODO:
      return state;
    case SAVE_TODO:
      return state;
    default:
      return state;
  }
}

export default todoList;
