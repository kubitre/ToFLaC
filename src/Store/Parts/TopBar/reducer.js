import { ADD_NEW_FILE, CLOSE_FILE, SET_ACTIVE_TAB } from "./constants";

const initState = {
    openedFiles: [
        {
            id: 1,
            name: "unknown"
        }
    ],
    active_file: 1,
}

export default function TBS_topBar(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case ADD_NEW_FILE:
            let openFiles = state.openedFiles;
            openFiles.push({id:payload.id, name: payload.name})
            return {...state, openedFiles: openFiles, active_file: payload.id}
        case CLOSE_FILE:
            let opened = []
            let active = 0
            console.log("state opened files changed to : ", payload)
            for(let i = 0; i < state.openedFiles.length; i ++){
                console.log("Iteration: state id: ", state.openedFiles[i].id, "")
                if (!(state.openedFiles[i].id === Number(payload))){
                    console.log("state id type: ", typeof(state.openedFiles[i].id), " no equals to: type ", typeof(Number(payload)))
                    opened.push(state.openedFiles[i])
                    active = state.openedFiles[i].id
                }
            }
            console.log("opened files after check iterations: ", opened)
            return {...state, openedFiles: opened, active_file: active}
        case SET_ACTIVE_TAB:
            return {...state, active_file: Number(payload)}
        default:
            return state;
    }
}