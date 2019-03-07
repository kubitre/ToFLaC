import React, {Component} from 'react';
import LeftPanel from '../../Components/LeftPanel';
import TabPanel from '../../Components/TabPanel';
import StatusBar from '../../Components/StatusBar';
import InputBlock from '../../Components/InputBlock';
import OutputBlock from '../../Components/OutputBlock';
import EmptyField from '../../Components/EmptyField';
import LinesShower from '../../Components/LinesShower';

// redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as userActions from '../../Store/Actions/user';
import * as windowActions from '../../Store/Actions/window';
import * as caretActions from '../../Store/Actions/caret';
import * as textActions from '../../Store/Actions/inputblock';

//styles
import './style.scss';


class MainPageComponent extends Component{
    constructor(props){
        super(props);
        this.renderBlock = this.renderBlock.bind(this);
    }
    renderBlock = (acrive_tab) => {
        switch(acrive_tab){
            case 0:
                return  <InputBlock >
                            <LinesShower />
                        </InputBlock>
            case 1:
                return  <OutputBlock >
                            <LinesShower/>
                        </OutputBlock>
            default: 
                return  <EmptyField />
        }
    }
    render = () => {
        return (
            <div className="main_page_container">
                <LeftPanel />
                <div className="main_block_" style={{width: '100%'}}>
                    <TabPanel />
                    {this.renderBlock(-1)}
                    <StatusBar/>
                </div>
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
        textEditorActions: bindActionCreators(textActions, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(MainPageComponent)