import { ADD_TODO, DELETE_TODO } from "./types"

//const BASE_URL = location.href.indexOf("localhost") > 0 ? "http://localhost:3000" : "";

export const addTodo = (todo) => async (
    dispatch
) => {
        dispatch({
            type: ADD_TODO,
            payload: todo
        })  
}

export const aTodo = (id) => async (
    dispatch
) => {
        dispatch({
            type: DELETE_TODO,
            payload: id
        })
    
}