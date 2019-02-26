
import {combineReducers} from 'redux';
import userState from './user';
import caretState from './caret';
import windowState from './window';

export default combineReducers({
    userState,
    caretState,
    windowState,
})
