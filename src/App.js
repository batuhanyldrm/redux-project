import React, {useState} from 'react';
import './App.css';
import Button from '@mui/material/Button';
import AddTodoPopUp from './AddTodoPopUp';
import TodoList from './TodoList';


function App() {
  const [open, setOpen] = useState(false);

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div >
      <AddTodoPopUp
      open={open}
      handleClose={handleClose}
      />
      <Button variant="contained" color="primary" style={{margin:"5px"}} onClick={() => setOpen(true)}>
        ADD TODO
      </Button>
      <TodoList/>
    </div>
  );
}

export default App;
