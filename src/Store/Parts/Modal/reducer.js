import { OPEN_MODAL_WINDOW, CLOSE_MW, CHANGE_SOURCE_MW } from "./constants";

const initState = {
    head: "test",
    body: "<div class=`body`>There is a small text</div>",
    open: false
}

export default function MWI_modalWindow(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case OPEN_MODAL_WINDOW:
        case CLOSE_MW:
            return {...state, open: payload}
        case CHANGE_SOURCE_MW:
            return {...state, head: payload.head, body: payload.body}
        default: 
            return state;
    }
}