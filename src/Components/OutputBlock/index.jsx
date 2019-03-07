import React from 'react';

import './style.scss';

export default function OutputBlock (props) {
    console.log(props);

    const {text_with_analys} = props;

    // document.getElementsByClassName("field_for_out").innerHTML = text_with_analys;

    return (
        <div className="output_block_component">
            {props.children}
            <div className="field_for_out"/>
        </div>
    )
}