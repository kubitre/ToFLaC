import '../Constants/window';
import { CREATE_NEW_CODE_FILE_WINDOW, OPEN_FILE_FROM_LOCAL_DISK, HELPER_PROGRAMM, SAVE_AS_FILE_TO_LOCAL_DISK, ABOUT_PROGRAMM } from '../Constants/window';

const initState = {
    "current": "mainpage",
}

export default function windowState (state = initState, action) {
    switch(action.type){
        case CREATE_NEW_CODE_FILE_WINDOW:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        case OPEN_FILE_FROM_LOCAL_DISK:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        case HELPER_PROGRAMM:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        case SAVE_AS_FILE_TO_LOCAL_DISK:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        case ABOUT_PROGRAMM:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        default:
            return state;
    }
}