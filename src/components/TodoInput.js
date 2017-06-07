import React from "react";
import PropTypes from "prop-types";
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';

class TodoInput extends React.Component {
  handleClick() {
    const description = this.refs.description.getValue();
    this.refs.description.input.value = "";
    this.props.save(description);
  }

  render() {
    return (
      <div>
        <TextField
          floatingLabelText="description"
          ref="description"
          errorText={this.props.message}
        />
        <RaisedButton
          label="Add"
          primary={true}
          onTouchTap={() => this.handleClick()}
          disabled={this.props.isFetching}
        />
      </div>
    );
  }
}

TodoInput.propTypes = {
  isFetching: PropTypes.bool,
  message: PropTypes.string
};

export default TodoInput;
