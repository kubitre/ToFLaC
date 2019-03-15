import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import './style.scss';
import ContentEditable from 'react-contenteditable';


class InputBlock extends Component {
    constructor(props){
        super(props);

        this.handleInputText = this.handleInputText.bind(this);
        this.html = "<div class='line_input'><br></div>";
        // console.log"Params: ", props);
    }
    handleInputText = (event) => {
        var pattern_edit_text = /<div class='line_input'>([0-9a-zA-Z]{0,}|<br>)<\/div>/gm;
        var pattern_line = /<div>/g;
        var pattern_text = /[0-9a-zA-Z \/;_+=-]+/;

        let text = event.target.value  

        console.log("handled text: ", text);
        if (text == ""){
            text = "<div></div>";
        }

        let allMatches = text.match(pattern_edit_text)
        console.log("mathces: ", allMatches);
        // console.log("amount lines: ", allMatches.length)

        let matchDivs = text.match(pattern_line);
        console.log('divs: ', matchDivs);

        let txt = "";

        text = text.replace(pattern_line, "<div class='line_input'>");
        if (!text.includes("<div>")){
            txt = text.match(pattern_text);    
        }

        this.html = text;
        console.log("text before regular match: ", text,"\ntext after regular match: ", txt);
                // console.log("text after transformations: ", text);
        // this.props.caretActions.setNewLinePosition(allMatches.length)
        // console.log("current column: ", allMatches[allMatches.length - 1])
    }

    render = () => { 
        const {html} = this.props.store;

        return(
            <div className="input_block_component">
                {this.props.children}
                <ContentEditable 
                    className="input_stream_block" 
                    html={html}
                    onChange={this.handleInputText} />
            </div>
        )
    }
}



function mapStateToProps (state) {
    return {
        store: state.inputBlockState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(InputBlock)