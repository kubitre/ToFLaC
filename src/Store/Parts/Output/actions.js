import { ADD_OUTPUT_BODY, CHANGE_OUTPUT_TYPE } from "./constants";

export function addOutput(data){
    return{
        type: ADD_OUTPUT_BODY,
        payload: data
    }
}

export function changeOutputType(newType){
    return{
        type: CHANGE_OUTPUT_TYPE,
        payload: newType,
    }
}