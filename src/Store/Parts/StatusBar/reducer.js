import { ADD_WARNINGS_AND_ERRORS, ADD_LINE_COL } from "./constants";

const initState = {
    "lang": "c++",
    "warnings_amount": 0,
    "errors_amount": 0,
    "line": 1,
    "column": 1,
}

export default function SBI_statusBar(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case ADD_WARNINGS_AND_ERRORS:
            return {...state, warnings_amount: payload.warnings, errors_amount: payload.errors}
        case ADD_LINE_COL:
            return {...state, line: payload.line, column: payload.column}
        default:
            return state;
    }
}