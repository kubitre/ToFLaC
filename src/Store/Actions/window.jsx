import '../Constants/window';
import { PROJECT_PAGE_LOCATION, ABOUT_PROGRAMM, SAVE_AS_FILE_TO_LOCAL_DISK, OPEN_FILE_FROM_LOCAL_DISK, CREATE_NEW_CODE_FILE_WINDOW, HELPER_PROGRAMM } from '../Constants/window';

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

export function openOpenModalWindow(state) {
    return {
        type: OPEN_FILE_FROM_LOCAL_DISK,
        payload: state
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