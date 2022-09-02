import { ADD_TODO, DELETE_TODO } from "../actions/types";

const Reducer = (state = {}, action) => {
    switch (action.type) {
        case ADD_TODO:
            return {...state, todos: action.payload}
        case DELETE_TODO:
            const newState = [...state.todos];
            const removedTodo = newState.filter((todo) =>
            todo.id !== action.payload)
            return {...state, todos: removedTodo}
        default:
            return state
    }
};

export default Reducer;