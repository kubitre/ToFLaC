import '../Constants/window';
import { PROJECT_PAGE_LOCATION, ABOUT_PROGRAMM, SAVE_AS_FILE_TO_LOCAL_DISK, OPEN_FILE_FROM_LOCAL_DISK, CREATE_NEW_CODE_FILE_WINDOW, HELPER_PROGRAMM, STATE_MAIN_BLOCK } from '../Constants/window';
import { CreatePacket, postRequest } from '../Api/RequestBuilder';
import { API_URL, url_cw } from '../Api/main';
import { EXCEPTION, ADD_WARNINGS_AND_ERRORS_AMOUNTS } from '../Constants/statusbar';
import { SHOW_OUTPUT_BLOCK } from '../Constants/mainblock';

export function openProjectLocation(state) {
    return {
        type: PROJECT_PAGE_LOCATION,
        payload: state
    }
}

export function openAboutProgramModalWindow(state) {
    return {
        type: ABOUT_PROGRAMM,
        payload: state
    }
}

export function saveAsProgramModalWindow(state){
    return {
        type: SAVE_AS_FILE_TO_LOCAL_DISK,
        payload: state
    }
}

export function openOpenModalWindow() {
    return {
        type: OPEN_FILE_FROM_LOCAL_DISK,
        payload: true
    }
}

export function createNewModalWindow(state){
    return {
        type: CREATE_NEW_CODE_FILE_WINDOW,
        payload: state
    }
}

export function openHelperModalWindow(state){
    return {
        type: HELPER_PROGRAMM,
        payload: state
    }
}

export function setStateMainBlock(state_number){
    return {
        "type": STATE_MAIN_BLOCK,
        "payload": state_number,
    }
}

export function StartFetchingDataFromBackend(body){
    return dispatch => {
        let request_packet = CreatePacket(API_URL + url_cw, JSON.stringify({"text": body}))

        dispatch({
            "type": SHOW_OUTPUT_BLOCK,
            "payload": true
        })
        fetch(request_packet)
        .then(response => {
            if (!response.ok) {
                throw new Error()
            }
            console.log(response.status)
            return response.json()
        })
        .then(response => {
            dispatch({type: "FETCHING_SUCCESS", payload: response});
            dispatch({type: ADD_WARNINGS_AND_ERRORS_AMOUNTS, "payload": {
                warnings: response.LexerAnalysPart.Errors.length,
                errors: response.SyntaxAnalysPart.Errors.length,
            }})
            })
        .catch(err => console.log(err))
    }
}