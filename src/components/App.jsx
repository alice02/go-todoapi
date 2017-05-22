import React from 'react';
import TodoList from './TodoList';
import axios from 'axios';

export default class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {todos: []};
  }

  componentDidMount() {
    axios.get('/api/tasks').then((response) => {
      this.setState({todos: response.data.data.tasks});
    }).catch((response) => {
      console.log("Error", response);
    })
  }

  render() {
    return (
      <div>
        <TodoList todos={this.state.todos}/>
      </div>
    );
  }

}
