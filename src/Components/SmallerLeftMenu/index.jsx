import React, {Component} from 'react';
import ElementLeftMenu from '../ElementLeftMenu';

import { connect } from 'react-redux';

import './style.scss';

class LeftMenu extends Component{
    render = () => {
        const {elements} = this.props.leftMenuFull;
        return(
            <div className="panel_menu_component small">
                {elements.map((element, index)=>{
                    return (
                        <ElementLeftMenu key={index} element={element}/>
                    )
                })}
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        leftMenuFull: state.LMI_leftMenu
    }
}

function mapDispatchToProps(dispatch) {
    return {
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(LeftMenu)