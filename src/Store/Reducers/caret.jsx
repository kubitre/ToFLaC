import { SET_LINE_POSITION } from "../Constants/caret";

const initState = {
    "line": 1,
    "column": 1,
    "lines": [1],
}

export default function caretState (state = initState, action) {
    switch(action.type){
        case SET_LINE_POSITION:
            // console.log(action, state)
            return {...state, line: action.payload}
        default:
            return state;
    }
    // return state;
}