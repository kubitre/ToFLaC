import React from 'react';

import './style.scss';

export default StatusBar = () => {
    console.log(this.props);
    const {column, line} = this.props.caret;
    const {warnsAmount, errorsAmount} = this.props.status;
    const {language} = this.props;

    return (
        <div className="status_bar_component">
            <div className="left_side">
                <h4 className="language">lang: {language}</h4>
            </div>
            <div className="right_side">
                <div className="info_block">
                    <h4 className="errors">{errorsAmount} Errors</h4>
                    <span className="separator">|</span>
                    <h4 className="warnings">{warnsAmount} Warnings</h4>
                </div>
                <div className="caret_info">Line: {line} | Col: {column}</div>
            </div>
        </div>
    )
}