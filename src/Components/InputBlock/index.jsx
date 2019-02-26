import React from 'react';

export default InputBlock = () => {
    console.log(this.props);

    return (
        <div className="input_block_component">
            <div className="left_side_line_numbers" readOnly />
            <div className="input_stream_block" ></div>
        </div>
    )
}