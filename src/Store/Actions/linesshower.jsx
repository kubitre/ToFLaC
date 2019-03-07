import { ADD_NEW_LINE, SET_LINES_TO_SHOW, REMOVE_LINE } from "../Constants/linesshower";

export function addNewLineToLineShower(){
    return {
        "type": ADD_NEW_LINE,
        "payload": +1
    }
}

export function setLinesToLineShower(lines){
    return {
        "type": SET_LINES_TO_SHOW,
        "payload": lines
    }
}

export function removeLineFromLineShower(){
    return {
        "type": REMOVE_LINE,
        "payload": -1
    }
}