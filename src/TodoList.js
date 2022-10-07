import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { IconButton } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { fetchTodos, deleteTodo} from './actions/todoActions';
import EditTodoPopUp from './EditTodoPopUp';

function TodoList(props) {
    const {fetchTodos, todos, deleteTodo} = props;

    const [changeTodoPopUp, setChangeTodoPopUp] = useState(false)
    const handleChangeTodoPopUpClose = () => setChangeTodoPopUp(false)
    const [id, setId] = useState("")

    const handleEdit = (todoID) => {
      setId(todoID);
      setChangeTodoPopUp(true);
    }

    useEffect(() => {
      fetchTodos()
    }, [])
    

    return(
        <div>
          <EditTodoPopUp
            openn={changeTodoPopUp}
            handleClosee={handleChangeTodoPopUpClose}
            id={id}
          />
            <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Todo</TableCell>
            <TableCell align="right">Is Completed</TableCell>
            <TableCell align="right">Created At</TableCell>
            <TableCell align="right">Updated At</TableCell>
            <TableCell align="right">Edit</TableCell>
            <TableCell align="right">Delete</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
         {todos.allTodos && todos.allTodos.map((item) => (
            <TableRow
              key={item.name}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
                {console.log(todos.allTodos,"aaaa")}
              <TableCell component="th" scope="row">
                {item.name}
              </TableCell>
              <TableCell align="right">{item.isCompleted}</TableCell>
              <TableCell align="right">{item.createdAt}</TableCell>
              <TableCell align="right">{item.updatedAt}</TableCell>
              <TableCell align="right">
                <IconButton
                  onClick={() => handleEdit(item.id)}
                >
                  <EditIcon size="small"/>
                </IconButton>
              </TableCell>
              <TableCell align="right">
                <IconButton
                  onClick={() => deleteTodo(item.id)}
                >
                  <DeleteIcon size="small"/>
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
        </div>
    )
}

const mapStateToProps = (state) => ({
    todos: state.todos
  });
  
  const mapDispatchToProps = (dispatch) => ({
    fetchTodos: () => {
      dispatch(fetchTodos());
    },
    deleteTodo: (id) => {
      dispatch(deleteTodo( id));
    },
  });

export default connect(mapStateToProps,mapDispatchToProps) (TodoList);