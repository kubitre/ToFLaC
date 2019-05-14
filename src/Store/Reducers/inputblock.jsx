import {COPY_TEXT_INPUT_BLOCK, ADD_TEXT_INPUT_BLOCK, PASTE_TEXT_INPUT_BLOCK, CUT_TEXT_INPUT_BLOCK, FLUSH_TEXT_INPUT_BLOCK} from '../Constants/inputblock';

const initState = {
    "code": "",
};

export default function inputBlockState(state = initState, action){
    switch(action.type){
        case COPY_TEXT_INPUT_BLOCK: 
        case PASTE_TEXT_INPUT_BLOCK:
        case CUT_TEXT_INPUT_BLOCK: 
        case ADD_TEXT_INPUT_BLOCK: 
        case FLUSH_TEXT_INPUT_BLOCK: 
            return {...state, 
                code: action.payload.code}
        default:
            return state;
    }
}