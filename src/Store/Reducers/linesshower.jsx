import { ADD_NEW_LINE, REMOVE_LINE, SET_LINES_TO_SHOW } from "../Constants/linesshower";

const initState = {
    "lines": [1]
}

export default function linesShowerState(state = initState, action){
    switch(action.type){
        case ADD_NEW_LINE:
            return {...state, lines: state.lines.push(state.lines[state.lines.length - 1] + 1)}
        case REMOVE_LINE:
            return {...state, lines: state.lines.pop()}
        case SET_LINES_TO_SHOW:
            return {...state, lines: action.payload}
        default:
            return state
    }
}