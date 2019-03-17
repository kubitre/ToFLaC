import { SET_FILE_ACTIVE, CLOSE_FILE, NEW_FILE } from "../Constants/toppanel";

export function setActiveFile(id){
    return {
        "type": SET_FILE_ACTIVE,
        "payload": id
    }
}

export function closeFile(id){
    return{
        "type": CLOSE_FILE,
        "payload": id,
    }
}

export function openNewFile(name){
    return{
        "type": NEW_FILE,
        "payload": {
            "filename": name
        }
    }
}