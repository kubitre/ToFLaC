import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as statusActions from '../../Store/Actions/statusbar';

import './style.scss';

class BlockHoverInformation extends Component{
    render = () => {
        console.log(this.props);
        const {errors_amount, warnings_amount, errors, warnings} = this.props.statusStore;
        return (
            <div className="block_hover" 
                onMouseEnter={() => this.props.statusActions.setHoverPanel()}
                onMouseLeave={() => this.props.statusActions.setClosePanel()}>

                <div className="errors_list">
                    <div className="header" style={{borderBottom: '1px solid #e6e6e6'}}>Ошибки: {errors_amount}</div>
                    <div className="list" style={{display: 'flex', flexDirection: 'column'}}>{errors.map((error, index) => 
                            <div className="error_li" key={error + index}>{error}</div>
                        
                        )}</div>
                </div>
                <div className="warnings_list"></div>
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        statusStore: state.statusBarState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        statusActions: bindActionCreators(statusActions, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(BlockHoverInformation)