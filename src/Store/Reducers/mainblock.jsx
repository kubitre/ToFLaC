import { SHOW_OUTPUT_BLOCK, CLOSE_OUTPUT_BLOCK } from "../Constants/mainblock";

const initState = {
    "output_block": false,
    "output_data": null
}

export default function outputBlockState(state = initState, action){
    switch(action.type){
        case SHOW_OUTPUT_BLOCK: 
            return {...state, output_block: action.payload}
        case CLOSE_OUTPUT_BLOCK:
            return {...state, output_block: action.payload}
        case "FETCHING_SUCCESS":
            return {...state, output_block: true, output_data: action.payload}
        default:
            return state;
    }
}