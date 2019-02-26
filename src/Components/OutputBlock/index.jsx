import React from 'react';

import './style.scss';

export default OutputBlock = () => {
    console.log(this.props);

    const {text_with_analys} = this.props;

    document.getElementsByClassName("field_for_out").innerHTML = text_with_analys;

    return (
        <div className="output_block_component">
            <p className="field_for_out"/>
        </div>
    )
}