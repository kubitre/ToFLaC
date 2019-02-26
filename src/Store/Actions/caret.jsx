import '../Constants/caret';
import { GET_CARET_POSITION, SET_COLUMN_POSITION, ADD_LINE_POSITION, REMOVE_LINE_POSITION, REMOVE_COLUMN_POSITION } from '../Constants/caret';

export function getCurrentCaretPosition(position){
    return {
        "type": GET_CARET_POSITION,
        "payload": position
    }
}

export function setNewColumnPosition(newposition){
    return{
        "type": SET_COLUMN_POSITION,
        "payload": newposition
    }
}

export function removeColumnPosition(position){
    return {
        "type": REMOVE_COLUMN_POSITION,
        "payload": position
    }
}

export function addNewLinePosition(newposition){
    return {
        "type": ADD_LINE_POSITION,
        "payload": newposition,
    }
}

export function removeLinePosition(position){
    return {
        "type": REMOVE_LINE_POSITION,
        "payload": position
    }
}