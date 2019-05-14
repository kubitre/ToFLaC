import React, { Component } from 'react';

import { connect } from 'react-redux';

import './style.scss';
import { bindActionCreators } from 'redux';

// import * as methods from '../../Store/Actions/window';

import * as fetching from '../../Store/Parts/LeftMenu/actions';
import * as Tab from '../../Store/Parts/TopBar/actions';

class ElementLeftMenu extends Component {

    handleLeftAction(event) {
        let action = event.target.getAttribute("action");

        switch (action) {
            case "create":
                console.log(this.props)
                let elemId = this.props.topBar.openedFiles.length + 1
                this.props.tabs.OpenNewFile("new", elemId)
                return

            case "execute":
                this.props.fetch.StartFetchingData(this.props.store.code)
                return

            default:
                return
        }
    }

    render = () => {
        const { avatar, id, action } = this.props.element;
        return (
            <div className="element_left_panel_menu" onClick={this.handleLeftAction.bind(this)} action={action}>
                <div className="img_element"
                    style={{
                        background: `url(${avatar})`,
                        backgroundSize: 'cover',
                    }}
                />
            </div>
        )
    }
}
function mapStateToProps(state) {
    return {
        store: state.ISI_inputState,
        topBar: state.TBI_topBar,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        fetch: bindActionCreators(fetching, dispatch),
        tabs: bindActionCreators(Tab, dispatch)
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(ElementLeftMenu)