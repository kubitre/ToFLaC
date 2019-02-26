import React, {Component} from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as userActions from '../../Store/Actions/user';
import * as windowActions from '../../Store/Actions/window';
import * as caretActions from '../../Store/Actions/caret';
import * as textEditorActions from '../../Store/Actions/textEditor';

class MainPageComponent extends Component{
    render = () => {
        return (
            <div className="ma in_page_container" style={{display: 'flex', flexDirection: 'column'}}>
            
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