import '../Constants/user';
import { GET_USER, LOGOUT_USER, GET_DIR_USER } from '../Constants/user';

export function getUserInfo(token){
    return {
        "type": GET_USER,
        "payload": token
    }
}

export function logoutUser(token){
    return {
        "type": LOGOUT_USER,
        "payload": token
    }
}

export function getProjects(token){
    return {
        "type": GET_DIR_USER,
        "payload": token
    }
}