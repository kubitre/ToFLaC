import { NEW_FILE, CLOSE_FILE, SET_FILE_ACTIVE} from "../Constants/toppanel";

const initState = {
    "files": [
        // {
        //     "id": 0,
        //     "filename": "test1.cpp"
        // },
        // {
        //     "id": 1,
        //     "filename": "test2.cpp"
        // }
    ],
    "active_file": 0
}

export default function topPanelState(state=initState, action){
    const{type, payload} = action;

    switch(type){
        case NEW_FILE:
            return {...state, files: state.files.push({id: payload.id, filename: payload.filename}), active_file: payload.id}
        case CLOSE_FILE:
            console.log(state);
            return {...state, files: state.files.splice(state.files.indexOf(payload.id), 1), active_file: payload.id === state.active_file ? state.files[state.files.length] : state.active_file}
        case SET_FILE_ACTIVE:
            return {...state, active_file: payload.id}
        default:
            return state;
    }
}

