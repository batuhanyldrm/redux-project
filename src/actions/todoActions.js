import { getTodos, removeTodos } from "../api/todoApi"
import { ADD_TODO, FETCH_TODO, DELETE_TODO } from "./types"

//const BASE_URL = location.href.indexOf("localhost") > 0 ? "http://localhost:3000" : "";

export const fetchTodos = (id) => async (
    dispatch
) => {
    const resp = await getTodos(id)
        dispatch({
            type: FETCH_TODO,
            payload: resp.data
        })
    
}

export const addTodo = (todo) => async (
    dispatch
) => {
        dispatch({
            type: ADD_TODO,
            payload: todo
        })  
}

export const deleteTodo = (id) => async (
    dispatch
) => {
    const resp = await removeTodos(id)
        dispatch({
            type: DELETE_TODO,
            payload: resp.id
        })
}