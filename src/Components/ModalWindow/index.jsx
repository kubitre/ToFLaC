import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as modalActions from '../../Store/Actions/modalwindow';

import './style.scss';

class ModalWindow extends Component{
    render = () => {
        const {header} = this.props.store;

        return (
            <div className="modal_window">
                <div className="header_part" style={{display: 'flex', flexDirection: 'row'}}>
                    <div className="title">{header}</div>
                    <span className="closing_file" onMouseDown={()=> this.props.modalActions.closeModalWindwow()}></span>
                </div>

            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        store: state.modalWindowState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        modalActions: bindActionCreators(modalActions, dispatch),
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(ModalWindow)