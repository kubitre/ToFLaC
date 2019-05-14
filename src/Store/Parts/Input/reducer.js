import { CHANGE_SOURCE_TEXT, ADD_TO_CASH, CHANGE_CONTEXT } from "./constants";

const initState = {
    code: " ",
    codes: [
        {
            "id": 1,
            "code": "",
        }
    ]
}

export default function ISI_inputState(state = initState, action){
    const {type, payload} = action;
    console.log()

    switch(type){
        case ADD_TO_CASH:
            let codes = state.codes
            codes.push(payload)
            return {...state, codes: codes}
        case CHANGE_SOURCE_TEXT:
            return {...state, code: payload}
        case CHANGE_CONTEXT:
            for(let i =0 ; i < state.codes.length; i ++){
                if (state.codes[i].id === Number(payload)){
                    return {...state, code: state.codes[i].code}
                }
            }
            return {...state}
        default:
            return state
    }
}