const axios = require("axios").default;

const getTodos = async () => {
    const resp = await axios.get("http://localhost:3001/todos")
}

const postTodos = async () => {
    const resp = await axios.post("http://localhost:3001/todos")
}