import { ADD_OUTPUT_BODY, CHANGE_OUTPUT_TYPE } from "./constants";

const initState = {
    type: -1, // 0 - tokens\ 1 - errors on lexer/ 2 - errors on syntax/ -1 -no choosed
    output: {
        LexerAnalysPart: {
            Tokens: null,
            Errors: null,
        },
        SyntaxAnalysPart: {
            Warnings: null,
            Errors: null,
            repair: null,
        },
    }
}

export default function OSI_outputState(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case ADD_OUTPUT_BODY:
            return {...state, output: payload}        
        case CHANGE_OUTPUT_TYPE:
            return {...state, type: Number(payload)}
        default:
            return state;
    }
}