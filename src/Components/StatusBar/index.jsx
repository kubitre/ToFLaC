import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as statusActions from '../../Store/Actions/statusbar';

import './style.scss';

class StatusBar extends Component{
    constructor(props){
        super(props);
    }


    render = () => {
        const {lang, warnings_amount, errors_amount} = this.props.store;
        const {line, col} = this.props.caret;
        const {state_main_block} = this.props.window;

        console.log(this.props);
        // const {}
        return (
            state_main_block == -1 ?

            null
            :
            <div className="status_bar_component">
                <div className="left_side">
                    <div className="language">lang: {lang}</div>
                </div>
                <div className="right_side">
                    <div className="info_block"
                         onMouseEnter={() => this.props.statusActions.setHoverPanel()}
                         onMouseLeave={() => setTimeout(this.props.statusActions.setClosePanel(), 500)}
                         >
                        <div className="errors" >{errors_amount} Errors&#8195;</div>
                        <span className="separator">|</span>
                        <div className="warnings">&#8195;{warnings_amount} Warnings</div>
                    </div>
                    <div className="caret_info">Line: {line} | Col: {col}</div>
                </div>
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
        store : state.statusBarState,
        caret: state.caretState,
        window: state.windowState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        statusActions: bindActionCreators(statusActions, dispatch)   
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(StatusBar)