import { ADD_TODO, DELETE_TODO, FETCH_TODO, UPDATE_TODO } from "../actions/types";

const Reducer = (state = {}, action) => {
    switch (action.type) {
        case ADD_TODO:
            return {...state, allTodos: [...state.allTodos, action.payload]}
        case FETCH_TODO:
            return {...state, allTodos: action.payload}
        case DELETE_TODO:
            return {...state, allTodos: [...state.allTodos].filter((item) => item !== action.payload)}
        case UPDATE_TODO:
            const editTodos = {...state};
            editTodos.allTodos.map((item, index) => {
                if(action.payload.id == item.id) {
                    editTodos.allTodos[index].name = action.payload.name
                }
            })
            return{...state, allTodos: editTodos.allTodos}
        default:
            return state
    }
};

export default Reducer;