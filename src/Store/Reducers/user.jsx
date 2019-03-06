import '../Constants/user';
import { GET_USER, GET_DIR_USER, LOGOUT_USER, CREATE_NEW_USER } from '../Constants/user';

const initState = {
    "username": "unknown",
    "token": null,
    "projects": null,
    "avatar": ""
};

export default function userState (state = initState, action) {
    switch(action.type){
        case GET_USER: 
            return {...state, packet: {
                type: action.type,
                payload: action.payload
                }
            }

        case GET_DIR_USER:
            return {...state, packet: {
                    type: action.type,
                    payload: action.payload
            }}

        case LOGOUT_USER:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}

        case CREATE_NEW_USER:
            return {...state, packet: {
                type: action.type,
                payload: action.payload
            }}
        default:
            return state;
    }
}