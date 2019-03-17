import '../Constants/window';
import { CREATE_NEW_CODE_FILE_WINDOW, OPEN_FILE_FROM_LOCAL_DISK, HELPER_PROGRAMM, SAVE_AS_FILE_TO_LOCAL_DISK, ABOUT_PROGRAMM, STATE_MAIN_BLOCK } from '../Constants/window';

const initState = {
    "current": "mainpage",
    "left_menu": [
        {
            "img": "#",
            "name": "test1",
            "visible_image": false
        },
        {
            "img": "#",
            "name": "test2",
            "visible_image": false
        },
        {
            "img": "#",
            "name": "test3",
            "visible_image": false
        }
    ],
    "state_main_block": 0,
}

export default function windowState (state = initState, action) {
    const {type, payload} = action;
    
    switch(type){
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
        case STATE_MAIN_BLOCK:
            return {...state, 
                state_main_block: payload
            }
        default:
            return state;
    }
}