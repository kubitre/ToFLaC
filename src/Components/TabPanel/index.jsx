import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as FileActions from '../../Store/Parts/TopBar/actions';
import * as InputActions from '../../Store/Parts/Input/actions';

import './style.scss';

class TabPanel extends Component {
    handleClick(event){
        let ident = event.target.getAttribute("ident");
        if (ident === null){
            let closer = event.target.getAttribute("closer")
            this.props.fileActions.CloseFile((closer));
            this.props.inputActions.ChangeContext(closer - 1);
        } else {
            this.props.fileActions.SetActiveTab(ident);
            this.props.inputActions.ChangeContext(ident);
        }
    }
    render = () => {
        console.log(this.props)
        const {openedFiles, active_file} = this.props.store;

        return (
            openedFiles.length === 0 ? 
            null:

            <div className="tab_panel_component">
                {
                    openedFiles.map((item, index) => 
                        <div className={active_file === item.id ? "element_tab _active": "element_tab"} onClick={this.handleClick.bind(this)} key={index} ident={item.id}>
                            <div className="text_name" ident={item.id}>{item.name}</div>
                            <span className="closing_file" closer={item.id}></span>
                        </div>    
                    )
                }
{/*                 
                 */}
                
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        store: state.TBI_topBar,
        inputSt: state.ISI_inputState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        fileActions: bindActionCreators(FileActions, dispatch),
        inputActions: bindActionCreators(InputActions, dispatch)
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(TabPanel)