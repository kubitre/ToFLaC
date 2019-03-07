import { SHOW_OUTPUT_BLOCK, CLOSE_OUTPUT_BLOCK } from "../Constants/mainblock";

export function showOutputBlock(){
    return {
        "type": SHOW_OUTPUT_BLOCK,
        "paylaod": true,
    }
}

export function closeOutputBlock(){
    return{
        "type": CLOSE_OUTPUT_BLOCK,
        "payload": false,
    }
}
