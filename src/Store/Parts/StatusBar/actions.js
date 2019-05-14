import { ADD_WARNINGS_AND_ERRORS, ADD_LINE_COL } from "./constants";

export function AddWarningsErrors(payload){
    return {
        type: ADD_WARNINGS_AND_ERRORS,
        payload: payload
    }
}

export function AddLineCol(data){
    return {
        type: ADD_LINE_COL,
        payload: {
            line: data.line,
            column: data.column
        }
    }
}