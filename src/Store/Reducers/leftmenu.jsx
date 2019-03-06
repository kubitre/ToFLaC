import {FULL_LEFT_MENU_CLOSE, FULL_LEFT_MENU_OPEN} from '../Constants/leftmenu';

const initState = {
    "full": false,
};

export default function fullLeftMenuState(state = initState, action){
    switch(action.type){
        case FULL_LEFT_MENU_OPEN: 
            return {...state, full: action.payload}
        case FULL_LEFT_MENU_CLOSE:
            return {...state, full: action.payload}
        default:
            return state;
    }
}