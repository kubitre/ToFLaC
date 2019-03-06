import { COPY_TEXT, CUT_TEXT, PASTE_TEXT } from '../Constants/text';

export function copyText(text){
    return {
        "type": COPY_TEXT,
        "payload": text
    }
}

export function cutText(text){
    return {
        "type": CUT_TEXT,
        "paylaod": text
    }
}

export function pasteText(text){
    return {
        "type": PASTE_TEXT,
        "payload": text
    }
}