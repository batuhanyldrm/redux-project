import { connect } from 'react-redux';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Button from '@mui/material/Button';
import { addTodo } from './actions/todoActions';
import { getTodos } from './api/todoApi';


function AddTodoPopUp(props) {
  
  const {open, handleClose, addTodo} = props;

  const handleCreateTodo = async (todo) => {
    await getTodos(todo
    ).then((res) => {
      addTodo(res)
    })
  }

    return(
        <>
        <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Add Todo</DialogTitle>
        <DialogContent>
          <DialogContentText>
            You can add todo
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Name"
            type="name"
            fullWidth
            variant="standard"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={() =>handleCreateTodo()}>ADD</Button>
        </DialogActions>
      </Dialog>
      </>
    );
}

const mapStateToProps = (state) => ({
});

const mapDispatchToProps = (dispatch) => ({
  addTodo: (todo) => {dispatch(addTodo(todo));},
  });

export default connect(mapStateToProps, mapDispatchToProps) (AddTodoPopUp);