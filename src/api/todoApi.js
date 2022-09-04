const axios = require("axios").default;

const getTodos = async () => {
    const resp = await axios.get("http://localhost:3001/todos")
    return resp;
}

const postTodos = async (data) => {
    const resp = await axios.post("http://localhost:3001/todos",{name:data})
    return resp;
}

const removeTodos = async (id) => {
    const resp = await axios.delete(`http://localhost:3001/todos/${id}`)
    return resp;
}

const updateTodos = async (id, data) => {
    const resp = await axios.put(`http://localhost:3001/todos/${id}`,{name:data})
    return resp;
}

module.exports = {
    getTodos,
    postTodos,
    removeTodos,
    updateTodos
}