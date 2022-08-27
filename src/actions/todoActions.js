import { INSERT_TODO, DELETE_TODO } from "./types"

export const insertTodo = (todo) => {
    return(dispatch) => {
        dispatch({
            type: INSERT_TODO,
            payload: todo
        })
    }
}

export const deleteTodo = (id) => {
    return(dispatch) => {
        dispatch({
            type: DELETE_TODO,
            payload: id
        })
    }
}