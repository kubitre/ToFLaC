import React, {Component} from 'react';

import './style.scss';
import ContentEditable from 'react-contenteditable';

export default class InputBlock extends Component {
    constructor(props){
        super(props);

        this.handleInputText = this.handleInputText.bind(this);
        console.log("Params: ", props);
    }
    handleInputText = (event) => {
        var pattern_edit_text = /<div>([0-9a-zA-Z]{0,}|<br>)<\/div>/gm;
        let text = event.target.value  
        console.log("handled text: ", text);
        if (text === null){

        }
        let allMatches = text.match(pattern_edit_text)
        console.log("mathces: ", allMatches);
        console.log("amount lines: ", allMatches.length)
        this.props.caretActions.setNewLinePosition(allMatches.length)
        // console.log("current column: ", allMatches[allMatches.length - 1])
    }

    render = () => { 
        const {lines, textEditor} = this.props 
        return(
            <div className="input_block_component">
                <div className="left_side_line_numbers">
                    {lines.map((line, index) => {
                        return(
                            <div className="element_left_side" key={index}>{line}</div>
                        )
                    })}
                </div>
                <ContentEditable 
                            className="input_stream_block" 
                            html={textEditor}
                            onChange={this.handleInputText} />
            </div>
        )
    }
}