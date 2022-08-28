import React, {useState} from 'react';
import './App.css';
import Button from '@mui/material/Button';
import AddTodoPopUp from './AddTodoPopUp';



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
      <Button variant="contained" color="primary" onClick={() => setOpen(true)}>
        ADD TODO
      </Button>
    </div>
  );
}

export default App;
