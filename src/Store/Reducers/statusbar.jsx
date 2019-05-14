import {STATUS_HOVER, STATUS_CLOSE, ADD_WARNINGS_AND_ERRORS_AMOUNTS} from '../Constants/statusbar';

const initState = {
    "lang": "c++",
    "warnings_amount": 0,
    "errors_amount": 0,
    "status_hover": false,
    "errors": [
        "строка 15, не указан тип"
    ],
    "warnings": [],
};

export default function statusBarState(state = initState, action){
    const {type, payload} = action;

    switch(type){
        case STATUS_HOVER:
        case STATUS_CLOSE:
            return {...state, status_hover: payload.status_hover}
        
        case ADD_WARNINGS_AND_ERRORS_AMOUNTS:
            return {...state, errors_amount: payload.errors, warnings_amount: payload.warnings}
        default:
            return state;
    }
}