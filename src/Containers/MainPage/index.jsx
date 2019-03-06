import React, {Component} from 'react';
import MainBlock from '../../Components/MainBlock';
import LeftPanel from '../../Components/LeftPanel';

// redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as userActions from '../../Store/Actions/user';
import * as windowActions from '../../Store/Actions/window';
import * as caretActions from '../../Store/Actions/caret';
import * as textEditorActions from '../../Store/Actions/textEditor';

//styles
import './style.scss';


class MainPageComponent extends Component{
    render = () => {
        const {window, caret, textEditor, textEditorActions, caretActions, windowActions} = this.props
        var elements_menu = [
            {
                id: 1,
                "name": "File",
                "img": "https://cdn0.iconfinder.com/data/icons/thin-files-documents/57/thin-071_file_document_code_binary-512.png",
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
                "img": "#",
                "visible_img": true
            }
        ]
        return (
            <div className="main_page_container">
                <LeftPanel elements_menu={elements_menu}/>
                <MainBlock 
                        window={window} 
                        caret={caret} 
                        textEditor={textEditor} 
                        textEditorActions={textEditorActions} 
                        caretActions={caretActions} 
                        windowActions={windowActions}/>
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        user: state.userState,
        caret: state.caretState,
        window: state.windowState,
        textEditor: state.textEditorState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        userActions: bindActionCreators(userActions, dispatch),
        windowActions: bindActionCreators(windowActions, dispatch),
        caretActions: bindActionCreators(caretActions, dispatch),
        textEditorActions: bindActionCreators(textEditorActions, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(MainPageComponent)