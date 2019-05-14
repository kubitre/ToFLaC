
import {combineReducers} from 'redux';

import TMI_topMenu from '../Parts/TopMenu/reducer';
import TBI_topBar from '../Parts/TopBar/reducer';
import SBI_statusBar from '../Parts/StatusBar/reducer';
import OSI_outputState from '../Parts/Output/reducer';
import MWI_modalWindow from '../Parts/Modal/reducer';
import LMI_leftMenu from '../Parts/LeftMenu/reducer';
import ISI_inputState from '../Parts/Input/reducer';

export default combineReducers({
    ISI_inputState,
    LMI_leftMenu,
    MWI_modalWindow,
    OSI_outputState,
    SBI_statusBar,
    TBI_topBar,
    TMI_topMenu,
    // userState,
    // caretState,
    // windowState,
    // analysState,
    // inputBlockState,
    // fullLeftMenuState,
    // outputBlockState,
    // linesShowerState,
    // statusBarState,
    // topPanelState,
    // modalWindowState,
    
})
