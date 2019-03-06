import React, {Component} from 'react';

import './style.scss';

class StatusBar extends Component{
    constructor(props){
        super(props);
    }


    render = () => {
        const {language} = props;
        const {}
        return (
            <div className="status_bar_component">
                <div className="left_side">
                    <div className="language">lang: {language}</div>
                </div>
                <div className="right_side">
                    <div className="info_block">
                        <div className="errors">{errorsAmount} Errors</div>
                        <span className="separator">|</span>
                        <div className="warnings">{warnsAmount} Warnings</div>
                    </div>
                    <div className="caret_info">Line: {line} | Col: {column}</div>
                </div>
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
    }
}

function mapDispatchToProps(dispatch) {
    return {
        
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(StatusBar)