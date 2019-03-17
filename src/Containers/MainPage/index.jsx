import React, {Component} from 'react';
import TabPanel from '../../Components/TabPanel';
import StatusBar from '../../Components/StatusBar';
import InputBlock from '../../Components/InputBlock';
import OutputBlock from '../../Components/OutputBlock';
import EmptyField from '../../Components/EmptyField';
import LinesShower from '../../Components/LinesShower';
import LeftMenu from '../../Components/SmallerLeftMenu';

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
        console.log("Main block: ")
        console.log(this.props);

        const {state_main_block} = this.props.window;
        const {files} = this.props.topPanel;
        if (files.length == 0 && state_main_block != -1){
            this.props.windowActions.setStateMainBlock(-1);
        }
        return (
            <div className="main_page_container">
                <LeftMenu />
                <div className="main_block_">
                    <TabPanel />
                    {this.renderBlock(state_main_block)}
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
        topPanel: state.topPanelState,
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