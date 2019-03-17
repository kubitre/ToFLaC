import { STATUS_HOVER, STATUS_CLOSE } from "../Constants/statusbar";

export function setHoverPanel(){
    return {
        "type": STATUS_HOVER,
        "payload": {
            status_hover: true
        }
    }
}

export function setClosePanel(){
    return{
        "type": STATUS_CLOSE,
        "payload": {
            status_hover: false
        }
    }
}