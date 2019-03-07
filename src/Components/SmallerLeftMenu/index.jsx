import React, {Component} from 'react';
import ElementLeftMenu from '../ElementLeftMenu';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as leftMenuActions from '../../Store/Actions/leftmenu';

import './style.scss';

class LeftMenu extends Component{
    constructor(props){
        super(props);   

        this.handleMouseLeave = this.handleMouseLeave.bind(this);
        this.handleMouseOver = this.handleMouseOver.bind(this);
    }
    handleMouseOver = (event) => {
        this.props.leftMenuActions.fullLeftMenuOpen();
    }

    handleMouseLeave = (event) => {
        this.props.leftMenuActions.fullLeftMenuClose();
    }
    render = () => {
        const {elements_menu} = this.props.leftMenuFull;
        return(
            <div className="smaller_panel_menu_component" 
                onMouseEnter={this.handleMouseOver} 
                onMouseLeave={this.handleMouseLeave}>
                {elements_menu.map((element, index)=>{
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
        leftMenuFull: state.fullLeftMenuState
    }
}

function mapDispatchToProps(dispatch) {
    return {
        leftMenuActions: bindActionCreators(leftMenuActions, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(LeftMenu)