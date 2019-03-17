
import {combineReducers} from 'redux';
import userState from './user';
import caretState from './caret';
import windowState from './window';
import analysState from './analisys';
import inputBlockState from './inputblock';
import fullLeftMenuState from './leftmenu';
import outputBlockState from './mainblock';
import linesShowerState from './linesshower';
import statusBarState from './statusbar';
import topPanelState from './toppanel';
import modalWindowState from './modalwindow';

export default combineReducers({
    userState,
    caretState,
    windowState,
    analysState,
    inputBlockState,
    fullLeftMenuState,
    outputBlockState,
    linesShowerState,
    statusBarState,
    topPanelState,
    modalWindowState,
})
