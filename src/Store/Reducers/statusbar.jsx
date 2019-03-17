import {STATUS_HOVER, STATUS_CLOSE} from '../Constants/statusbar';

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
        
        default:
            return state;
    }
}