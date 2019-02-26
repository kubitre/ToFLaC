import '../Constants/text';
import { COPY_TEXT, PASTE_TEXT, CUT_TEXT } from '../Constants/text';

const initState = {
    "text": ""
};

export default function textEditorState(state = initState, action){
    switch(action.type){
        case COPY_TEXT: 
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        case PASTE_TEXT:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        case CUT_TEXT: 
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        default:
            return state;
    }
}