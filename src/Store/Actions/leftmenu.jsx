import { FULL_LEFT_MENU_OPEN, FULL_LEFT_MENU_CLOSE } from "../Constants/leftmenu";

export function fullLeftMenuOpen(){
    return {
        "type": FULL_LEFT_MENU_OPEN,
        "payload": true,
    }
}

export function fullLeftMenuClose(){
    return {
        "type": FULL_LEFT_MENU_CLOSE,
        "payload": false,
    }
}