
import {combineReducers} from 'redux';
import userState from './user';
import caretState from './caret';
import windowState from './window';
import analysState from './analisys';
import textEditorState from './textEditor';
import fullLeftMenuState from './leftmenu';

export default combineReducers({
    userState,
    caretState,
    windowState,
    analysState,
    textEditorState,
    fullLeftMenuState,
})
