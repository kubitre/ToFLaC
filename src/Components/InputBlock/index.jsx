import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import './style.scss';
import ContentEditable from 'react-contenteditable';


class InputBlock extends Component {
    constructor(props){
        super(props);

        this.handleInputText = this.handleInputText.bind(this);
        // console.log("Params: ", props);
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
        const textEditor = "<div><br></div>"
        return(
            <div className="input_block_component">
                {this.props.children}
                <ContentEditable 
                            className="input_stream_block" 
                            html={textEditor}
                            onChange={this.handleInputText} />
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

export default connect(mapStateToProps, mapDispatchToProps)(InputBlock)