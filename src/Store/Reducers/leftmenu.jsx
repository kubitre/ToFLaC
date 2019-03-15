import {FULL_LEFT_MENU_CLOSE, FULL_LEFT_MENU_OPEN} from '../Constants/leftmenu';

const initState = {
    "full": false,
    "elements_menu": [
        {
            id: 1,
            "name": "File",
            "img": "http://www.myiconfinder.com/uploads/iconsets/256-256-003ea7262e1af37558f37d061e90e4a8.png",
            "visible_img": true
        },
        {
            id: 2,
            "name": "Правка",
            "img": "https://image.flaticon.com/icons/svg/1373/1373015.svg",
            "visible_img": true,
        },
        {
            id: 3,
            "name": "Справка",
            "img": "https://support.apple.com/content/dam/edam/applecare/images/en_US/osx/featured-content-forgot-id-icon_2x.png",
            "visible_img": true
        }
    ]
};

export default function fullLeftMenuState(state = initState, action){
    switch(action.type){
        case FULL_LEFT_MENU_OPEN: 
            return {...state, full: action.payload}
        case FULL_LEFT_MENU_CLOSE:
            return {...state, full: action.payload}
        default:
            return state;
    }
}