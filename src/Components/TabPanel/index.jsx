import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import * as panelActions from '../../Store/Actions/toppanel';
import * as windowActions from '../../Store/Actions/window';

import './style.scss';

class TabPanel extends Component {
    render = () => {
        console.log(this.props.store)
        const {files, active_file} = this.props.store;
        return (
            files.length === 0 ? 
            null:

            <div className="tab_panel_component">
                {
                    files.map((item, index) => 
                        <div className={active_file === item.id ? "element_tab _active": "element_tab"} key={index}>
                            <div className="text_name">{item.filename}</div>
                            <span className="closing_file" onMouseDown={() => this.props.panelActions.closeFile(item.id)}></span>
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
        store: state.topPanelState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        panelActions: bindActionCreators(panelActions, dispatch),
        windowActions: bindActionCreators(windowActions, dispatch),

    }
}

export default connect(mapStateToProps, mapDispatchToProps)(TabPanel)