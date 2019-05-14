import { CHANGE_SOURCE_TEXT, ADD_TO_CASH, CHANGE_CONTEXT } from "./constants";

export function SetupSource(txt){
    return {
        type: CHANGE_SOURCE_TEXT,
        payload: txt,
    }
}

export function AddCodes(id, code){
    return {
        type: ADD_TO_CASH,
        payload: {
            id: id,
            code: code,
        }
    }
}

export function ChangeContext(id){
    return {
        type: CHANGE_CONTEXT,
        payload: id,
    }
}