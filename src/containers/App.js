import React from "react";
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';import Header from "./Header";
import TodoInput from "./TodoInput";
import TodoList from "./TodoList";

const App = () => {
  return (
    <MuiThemeProvider>
      <div>
        <Header />
        <TodoInput />
        <TodoList />
      </div>
    </MuiThemeProvider>
  );
}

export default App;
