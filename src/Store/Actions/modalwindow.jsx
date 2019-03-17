import { OPEN_MODAL_WINDOW, CLOSE_MODAL_WINDOW } from "../Constants/modalwindow";

export function openModalWindow(type, header){
    return {
        "type": OPEN_MODAL_WINDOW,
        "payload": {
            "type": type,
            "header": header,
        }
    }
}

export function closeModalWindwow(){
    return {
        "type": CLOSE_MODAL_WINDOW,
        "payload": null
    }
}