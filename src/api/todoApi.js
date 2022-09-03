const axios = require("axios").default;

const getTodos = async () => {
    const resp = await axios.get("http://localhost:3001/todos")
    return resp;
}

const postTodos = async () => {
    const resp = await axios.post("http://localhost:3001/todos")
    return resp;
}

const removeTodos = async (id) => {
    const resp = await axios.delete(`http://localhost:3001/todos/${id}`)
    return resp;
}

module.exports = {
    getTodos,
    postTodos,
    removeTodos
}