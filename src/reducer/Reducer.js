const Reducer = (state = 0, action) => {
    switch (action.type) {
        case "DEPOSİT":
            return state + action.payload;
        case "WITHDRAW":
            return state - action.payload;
        default:
            return state
    }
};

export default Reducer;