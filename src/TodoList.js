import React, { useEffect } from 'react';
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
import { fetchTodos, deleteTodo } from './actions/todoActions';

function TodoList(props) {
    const {fetchTodos, id, todos, deleteTodo} = props;

    useEffect(() => {
      fetchTodos(id)
    }, [])
    

    return(
        <div>
            <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Todo</TableCell>
            <TableCell align="right">Is Completed</TableCell>
            <TableCell align="right">Created At</TableCell>
            <TableCell align="right">Updated At</TableCell>
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
    //id: state.data.id
  });
  
  const mapDispatchToProps = (dispatch) => ({
    fetchTodos: (id) => {
      dispatch(fetchTodos( id));
    },
    deleteTodo: (id) => {
      dispatch(deleteTodo( id));
    },
  });

export default connect(mapStateToProps,mapDispatchToProps) (TodoList);