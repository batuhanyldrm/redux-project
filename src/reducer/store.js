import { createStore, applyMiddleware } from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import reducers from "./index";

function configureStore() {
  return createStore(reducers, composeWithDevTools(applyMiddleware()));
}

export default configureStore;
