import {OPEN_MODAL_WINDOW, CLOSE_MODAL_WINDOW} from '../Constants/modalwindow';

const initState = {
    "open": false,
    "type": 0, // 0 - open, 1 - save, 2 - new
    "header": "Открыть файл",
}

export default function modalWindowState(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case OPEN_MODAL_WINDOW:
            return {...state, open: true, type: payload.type, header: payload.header}
        case CLOSE_MODAL_WINDOW:
            return {...state, open: false}
        default:
            return state;
    }
}