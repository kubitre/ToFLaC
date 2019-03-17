import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as modalActions from '../../Store/Actions/modalwindow';

import './style.scss';
import ModalWindow from '../ModalWindow';

class EmptyField extends Component{
    render = () => {
        const {open} = this.props.modalStore;

        return (
            <div 
                className="empty_field" 
                style={{
                    background: "#d35400", 
                    width: '100%', 
                    height: '100%', 
                    display: 'flex', 
                    flexDirection: 'column', 
                    justifyContent: 'center', 
                    alignItems: 'center',
                    color: 'white'}}>
                   {open ? 
                        <ModalWindow />

                        :
                        <div className="open_modal_window" style={{marginTop: "45px"}}>
                            <h4>Вы можете </h4>
                            <div className="open_file_modal_window _button_open" onMouseDown={() => this.props.modalActions.openModalWindow(0, "Открыть файл")}>Открыть файл</div>
                            <div className="create_new_file_modal_window _button_open" onMouseDown={() => this.props.modalActions.openModalWindow(2, "Создать файл")}>Создать файл</div>
                        </div>
                    }
                     <h2>Вы пока не открыли ни одного файла</h2>
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
        store: state.inputBlockState,
        statusStore: state.statusBarState,
        modalStore: state.modalWindowState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        modalActions: bindActionCreators(modalActions, dispatch),
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(EmptyField)