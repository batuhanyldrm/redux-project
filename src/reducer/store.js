import { createStore } from "redux";
import reducers from "./index";
import Reducer from "./Reducer";

export const store = createStore(
    reducers,
    {}
)