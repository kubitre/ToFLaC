import React, {Component} from 'react';

import {connect} from 'react-redux';

import './style.scss';
import TopMenuElement from '../Element';

class TopMenu extends Component {
    render = () => {
        console.log(this.props)
        const {elements} = this.props.store;
        return (
            <div className="topmenu_container">
                {elements.map((element, index) => {
                    return (
                        <TopMenuElement key={index} elemen={element} />)
                })}
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
        store: state.TMI_topMenu,
    }
}

function mapDispatchToProps(dispatch) {
    return {
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(TopMenu)