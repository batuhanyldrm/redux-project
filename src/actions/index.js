export const depositMoney = (amount) => {
    return(dispatch) => {
        dispatch({
            type: "DEPOSIT",
            payload: amount
        })
    }
}

export const widthdrawMoney = (amount) => {
    return(dispatch) => {
        dispatch({
            type: "WITHDRAW",
            payload: amount
        })
    }
}