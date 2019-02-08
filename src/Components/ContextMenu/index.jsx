import React,{Component} from 'react';

import './style.css';

export default class ContextMenu extends Component{
    render = () => {
        return (
            <div className="dropdown">
                <button className="dropbtn">{this.props.name}</button>
                <div className="dropdown-content">
                    {this.props.children}
                </div>
            </div>
        )
    }
}