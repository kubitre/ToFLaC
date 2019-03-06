import {FULL_LEFT_MENU_CLOSE, FULL_LEFT_MENU_OPEN} from '../Constants/leftmenu';

const initState = {
    "lang": "None",
    "warnings": 0,
    "errors": 0,
    "errors_hover": false,
    "warnings_hover": false,
};

export default function statusBarState(state = initState, action){
    switch(action.type){
        case HOVER_WARNINGS: 
            return {...state, warnings_hover: action.payload}
        case CLOSE_WARNINGS:
            return {...state, warnings_hover: action.payload}
        case HOVER_ERRORS: 
            return {...state, errors_hover: action.payload}
        case CLOSE_ERRORS:
            return {...state, errors_hover: action.payload}
        case ADD_WARNING: 
            return {...state, warnings: action.payload}
        case REMOVE_WARNING:
            return {...state, warnings: action.payload}
        case ADD_ERROR:
            return {...state, errors: action.payload}
        case REMOVE_ERROR: 
            return {...state, errors: action.payload}
        default:
            return state;
    }
}