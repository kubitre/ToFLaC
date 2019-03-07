import {COPY_TEXT_INPUT_BLOCK, ADD_TEXT_INPUT_BLOCK, PASTE_TEXT_INPUT_BLOCK, CUT_TEXT_INPUT_BLOCK, FLUSH_TEXT_INPUT_BLOCK} from '../Constants/inputblock';

const initState = {
    "body_output": "<div><br></div>",
    "body_for_analys": "",
};

export default function inputBlockState(state = initState, action){
    switch(action.type){
        case COPY_TEXT_INPUT_BLOCK: 
        case PASTE_TEXT_INPUT_BLOCK:
        case CUT_TEXT_INPUT_BLOCK: 
        case ADD_TEXT_INPUT_BLOCK: 
        case FLUSH_TEXT_INPUT_BLOCK: 
            return {...state, 
                body_output: action.payload.body_output,
                body_for_analys: action.payload.body_for_analys}
        default:
            return state;
    }
}