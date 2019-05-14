import { CreatePacket } from "../../Api/RequestBuilder";
import { API_URL, url_cw } from "../../Api/main";
import { ADD_WARNINGS_AND_ERRORS } from "../StatusBar/constants";
import { ADD_OUTPUT_BODY } from "../Output/constants";

export function StartFetchingData(code) {
    return dispatch => {
        let request_packet = CreatePacket(API_URL + url_cw, JSON.stringify({"text": code}))

        fetch(request_packet)
        .then(response => {
            if (!response.ok) {
                throw new Error()
            }
            console.log(response.status)
            return response.json()
        })
        .then(response => {
            dispatch({type: ADD_OUTPUT_BODY, payload: response});
            dispatch({type: ADD_WARNINGS_AND_ERRORS, "payload": {
                warnings: response.LexerAnalysPart.Errors.length,
                errors: response.SyntaxAnalysPart.Errors.length,
            }})
            })
        .catch(err => console.log(err))
    }
}