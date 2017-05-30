import axios from "axios";

axios.defaults.baseURL = "http://localhost:1323";
axios.defaults.headers.post["Content-Type"] = "application/json"

export const FETCH_TODOS = "FETCH_TODOS";
export function fetchTodos() {
  return {
    type: FETCH_TODOS
  };
}

export const RECEIVE_TODOS = "RECEIVE_TODOS";
export function receiveTodos(todos) {
  return {
    type: RECEIVE_TODOS,
    todos: todos
  };
}

export const ADD_TODO = "ADD_TODO";
export function addTodo(todo) {
  return {
    type: ADD_TODO,
    todo: todo
  };
}

export const ADD_FAILED= "ADD_FAILED";
export function addFailed(message) {
  return {
    type: ADD_FAILED,
    message: message
  };
}

export const SAVE_TODO = "SAVE_TODO";
export function saveTodo(todo) {
  return {
    type: SAVE_TODO,
    todo: todo
  };
}

export function createTodo(description) {
  var todo = {completed: false, description: description};
  return (dispatch, getState) => {
    dispatch(addTodo(todo));
    axios.post("/api/tasks", todo)
         .then(response => {
           todo.id = response.data.data.id;
           dispatch(saveTodo(todo));
         }).catch(error => {
           dispatch(addFailed(error.response.data.data.info));
         });
  }
}

export const FINISH_TODO = "FINISH_TODO";
export function finishTodo(id) {
  return {
    type: FINISH_TODO,
    id: id
  };
}

function getTodos(store) {
  return dispatch => {
    dispatch(fetchTodos());
    return axios.get("/api/tasks")
    .then(response => {
      dispatch(receiveTodos(response.data.data.tasks));
    });
  };
}

export function getTodosIfNeeded() {
  return (dispatch, getState) => {
    if (getState().isFetching) {
      return Promise.resolve();
    } else {
      return dispatch(getTodos());
    }
  };
}
