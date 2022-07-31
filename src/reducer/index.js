import {createStore, combineReducers} from "redux";
import expensesReducer from "./Reducer";

function Store() {
    const store = createStore(
        combineReducers({
            expenses: expensesReducer,
        })
    );
    console.log(store.getState())
}

export default Store;