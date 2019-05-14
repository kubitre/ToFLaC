import React, {Component} from 'react';
import TabPanel from '../../Components/TabPanel';
import StatusBar from '../../Components/StatusBar';
import InputBlock from '../../Components/InputBlock';
import OutputBlock from '../../Components/OutputBlock';
import EmptyField from '../../Components/EmptyField';
import LinesShower from '../../Components/LinesShower';
import LeftMenu from '../../Components/SmallerLeftMenu';
import TopMenu from '../../Components/TopMenu/Entry';

// redux
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as userActions from '../../Store/Actions/user';
import * as windowActions from '../../Store/Actions/window';
import * as caretActions from '../../Store/Actions/caret';
import * as textActions from '../../Store/Actions/inputblock';

//styles
import './style.scss';
import ModalWindow from '../../Components/ModalWindow';


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
        // console.log("Main block: ")
        // console.log(this.props);

        // if (files.length == 0 && state_main_block != -1){
        //     this.props.windowActions.setStateMainBlock(-1);
        // }

        console.log("modal window state: ", this.props.storeModalWindow)

        const {open} = this.props.storeModalWindow;
        return (
            <div className="main_page_container">
                {open ? 
                    <ModalWindow />
                    :
                    null
                }
                <div className="vertical">
                    <TopMenu/>
                    <div className="horizontal">
                        <LeftMenu />
                        <div className="main_block_">
                            <TabPanel />
                            {this.renderBlock(0)}
                            {this.renderBlock(1)}
                            <StatusBar/>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        storeModalWindow: state.MWI_modalWindow
    }
}

function mapDispatchToProps(dispatch) {
    return {
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(MainPageComponent)