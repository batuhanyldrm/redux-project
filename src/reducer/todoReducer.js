import { ADD_TODO, FETCH_TODO } from "../actions/types";

const Reducer = (state = {}, action) => {
    switch (action.type) {
        case ADD_TODO:
            return {...state, todos: action.payload}
        case FETCH_TODO:
            return {...state, allTodos: action.payload}
        default:
            return state
    }
};

export default Reducer;