import { getTodos } from "../api/todoApi"
import { ADD_TODO, FETCH_TODO } from "./types"

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