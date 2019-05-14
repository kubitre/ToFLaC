import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as modalActions from '../../Store/Parts/Modal/actions';

import './style.scss';

class ModalWindow extends Component{
    render = () => {
        const {head, body} = this.props.store;

        return (
            <div className="modal_window">
                <div className="palette">
                    <div className="header_part" style={{display: 'flex', flexDirection: 'row'}}>
                        <div className="title">{head}</div>
                        <span className="closing_file" onMouseDown={()=> this.props.modalActions.CloseModalWindow()}></span>
                    </div>
                    <div className="body_part">
                        <div className="body" dangerouslySetInnerHTML={{__html:body}}/>
                    </div>
                </div>
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        store: state.MWI_modalWindow,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        modalActions: bindActionCreators(modalActions, dispatch),
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(ModalWindow)