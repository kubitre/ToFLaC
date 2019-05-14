import { ADD_NEW_FILE, CLOSE_FILE, SET_ACTIVE_TAB } from "./constants";

export function OpenNewFile(name, id){
    return {
        type: ADD_NEW_FILE,
        payload: {
            name: name,
            id: id
        }
    }
}

export function CloseFile(id){
    return {
        type: CLOSE_FILE,
        payload: id
    }
}

export function SetActiveTab(id){
    return {
        type: SET_ACTIVE_TAB,
        payload: id
    }
}