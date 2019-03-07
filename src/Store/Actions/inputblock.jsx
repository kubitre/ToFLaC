import { COPY_TEXT, CUT_TEXT, PASTE_TEXT, COPY_TEXT_INPUT_BLOCK, CUT_TEXT_INPUT_BLOCK, PASTE_TEXT_INPUT_BLOCK, ADD_TEXT_INPUT_BLOCK, FLUSH_TEXT_INPUT_BLOCK } from '../Constants/inputblock';

export function copyTextInputBlock(text){
    return {
        "type": COPY_TEXT_INPUT_BLOCK,
        "payload": {
            "body_output": text.body_output,
            "body_for_analys": text.body_for_analys,
        }
    }
}

export function cutTextInputBlock(text){
    return {
        "type": CUT_TEXT_INPUT_BLOCK,
        "payload": {
            "body_output": text.body_output,
            "body_for_analys": text.body_for_analys,
        }
    }
}

export function pasteTextInputBlock(text){
    return {
        "type": PASTE_TEXT_INPUT_BLOCK,
        "payload": {
            "body_output": text.body_output,
            "body_for_analys": text.body_for_analys,
        }
    }
}

export function addTextToInputBlock(text){
    return {
        "type": ADD_TEXT_INPUT_BLOCK,
        "payload": {
            "body_output": text.body_output,
            "body_for_analys": text.body_for_analys,
        }
    }
}

export function flushTextFromInputBlock(text){
    return {
        "type": FLUSH_TEXT_INPUT_BLOCK,
        "payload": {
            "body_output": text.body_output,
            "body_for_analys": text.body_for_analys,
        }
    }
}