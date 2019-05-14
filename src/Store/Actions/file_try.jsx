import { Examples } from "../Parts/Files/examples";
import { SetupSource, AddCodes } from "../Parts/Input/actions";
import { OpenNewFile } from "../Parts/TopBar/actions";

export function OpenFile(data, sourceCode){
    return dispatch => {
        dispatch(OpenNewFile(data.name, data.id))
        dispatch(SetupSource(sourceCode))
    }
}

export function OpenFileExample(param, currentId, currentTxt){
    return dispatch => {
        let example = Examples[param - 1]

        dispatch(OpenFile({
            id: example.id,
            name: example.name
        }, example.source))
        
    }
}

export function SetActiveTab(id){
    return dispatch => {
        dispatch(SetActiveTab(id))
    }
}