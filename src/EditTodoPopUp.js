import React, { useState, useEffect } from 'react'
import { connect } from 'react-redux';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import { updateTodo } from './actions/todoActions';
import { updateTodos } from './api/todoApi';
/* import TodoList from './TodoList'; */

function AddTodoPopUp(props) {
  
  const {openn, handleClosee, todos, updateTodo, id} = props;

  const [data, setData] = useState(todos.name)


  const handleChangeTodo = async () => {
    await updateTodos(id, data)
    .then(() => {
        updateTodo(data)
    })
  }

 /*  useEffect(() => {
    setData({name: todos.name})

  }, [todos]) */
  

    return(
        <>
        {/* <TodoList
            handleChangeTodo={handleChangeTodo}
        /> */}
        <Dialog open={openn} onClose={handleClosee}>
        <DialogTitle>Add Todo</DialogTitle>
        <DialogContent>
          <DialogContentText>
            You can add todo
          </DialogContentText>
          <TextField
            value={data}
            onChange={(e) => setData(e.target.value)}
            autoFocus
            margin="dense"
            id="data"
            label="Todo"
            type="text"
            fullWidth
            variant="standard"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClosee}>Cancel</Button>
          <Button onClick={() => handleChangeTodo()}>UPDATE</Button>
        </DialogActions>
      </Dialog>
      </>
    );
}

const mapStateToProps = (state) => ({
    todos: state.todos
});

const mapDispatchToProps = (dispatch) => ({
  updateTodo: (id,data) => {
    dispatch(updateTodo(id,data));
  },
  });

export default connect(mapStateToProps, mapDispatchToProps) (AddTodoPopUp);