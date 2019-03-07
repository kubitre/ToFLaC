import { SHOW_OUTPUT_BLOCK, CLOSE_OUTPUT_BLOCK } from "../Constants/mainblock";

const initState = {
    "output_block": false,
}

export default function outputBlockState(state = initState, action){
    switch(action.type){
        case SHOW_OUTPUT_BLOCK: 
            return {...state, output_block: action.payload}
        case CLOSE_OUTPUT_BLOCK:
            return {...state, output_block: action.payload}
        default:
            return state;
    }
}