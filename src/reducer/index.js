import {createStore, combineReducers} from "redux";
import accountReducer from "./Reducer";

const reducers = combineReducers({
    account: accountReducer
});

export default reducers;