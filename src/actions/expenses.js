
export const INITIAL_STATE ={
    expenses: [{
        id: "faruk",
        description: "çarşamba",
        note: "yarın cuma",
        amount: 65359,
        createdAt: 0
    }],
    filters: {
        text: "rent",
        sortBy: "amount", //date or amount
        startDate: undefined,
        endDate: undefined
    }
};

export const expensesReducer = (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case "REMOVE_EXPENSE":
            return state.filter(({id}) =>{
                return id !== action.id
            })
        default:
            return state;
    }
}

export const removeExpenses = ({ id } = {}) => ({
    type: "REMOVE_EXPENSE",
    id
})