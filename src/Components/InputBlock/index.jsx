import React, {Component} from 'react';
import {Button} from 'reactstrap';

import './index.scss';

export default class InputBlockComponent extends Component{
    constructor(props){
        super(props);
        this.state = {
            "amountLines": 1,
            "lines": "1",
            "text": "",
            "currentPositionCursor": {
                "line": 1,
                "col": 1
            },
            "selection": {
                "startoffset": 0,
                "endoffset": 0,
                "rangeAmount": 0,
            }
        }
        this.handleAddNeLine = this.handleAddNeLine.bind(this);
        this.handleBackspace = this.handleBackspace.bind(this);
        this.getCaretPosition = this.getCaretPosition.bind(this);
    }
    handleAddNeLine = (event) => {
        if (event.key === 'Enter') {
            this.setState({
                amountLines: this.state.amountLines + 1,
                lines: this.state.lines + "\n" + (this.state.amountLines + 1)
            });
        }
    }

    handleBackspace = (event) => {
        if(event.key === 'Backspace'){
            this.setState({
                amountLines: this.state.amountLines - 1 !== 0 ? this.state.amountLines - 1 : 1,
                lines: this.state.amountLines === 1 ? "1" : 
                    this.state.amountLines < 10 ? this.state.lines.substr(0, this.state.lines.length - 2) : 
                    this.state.amountLines < 100 ? this.state.lines.substr(0, this.state.lines.length - 3) : 
                    this.state.amountLines < 1000 ? this.state.lines.substr(0, this.state.lines.length - 4) : 5
            })
        }
    }

    handleChangeTextEdit = (event) => {
        const {keyCode} = event

        console.log("key upp: ", keyCode)
        
    }

    handleScroll = (event) => {
        document.getElementById("#line_panel").scrollTo(document.getElementById("#line_panel").scrollHeight, document.getElementById("#line_panel").scrollHeight);
    }

    getCaretPosition(event) {
        var editableDiv = event.target;
        console.log(editableDiv)
        var caretPos = 0,
          sel, range;
        if (window.getSelection) {
          sel = window.getSelection();
          console.log("sel: ", sel)
          var startSelection = 0
          var endSelection = 0
          if (sel.rangeCount) {
            range = sel.getRangeAt(0);
            console.log("range: ", range.endOffset)
            caretPos = range.endOffset
            startSelection = range.startOffset
            endSelection = range.endOffset
          }
          this.setState({
            currentPositionCursor: {
                "col": caretPos,
                "line": this.state.currentPositionCursor.line
            },
            selection: {
                "startoffset": startSelection,
                "endoffset": endSelection 
            }
           })
        } 
      }

    render = () => {
        return(
            <div className="inputblock_component" style={{display: 'flex', flexDirection: 'column', height: this.props.height, width: this.props.width, marginTop: 30}}>
                <div className="control_panel" style={{display: 'flex', flexDirection: 'row', height: window.innerHeight < 500 ? '15vh' : '5vh', marginLeft: 10, width: "100%"}}>
                    <Button color="success" style={{marginRight: '5px'}}>Execute</Button>
                    <Button color="danger">Stop</Button>
                </div>
                <div className="input_panel" style={{width: '100%', display: 'flex', flexDirection: 'row', height: this.props.height, paddingTop: '10px'}}>
                    <textarea id="#line_panel" className="code_block_input_area"
                        readOnly 
                        name="line_numbers" 
                        value={this.state.lines} 
                        style={{width: '70px', overflow: 'hidden', textAlign: 'center',
                                height: this.props.height        
                        }} 
                        onChange={this.handleChangeTextEdit}/>
                    <div className="code_block_input_area" id="editor" contentEditable={true} 
                        onKeyPress={this.handleAddNeLine}
                        onScroll={this.handleScroll}
                        onKeyDown={this.getCaretPosition}
                        onKeyUp={this.getCaretPosition}
                        placeholder="input your code here"
                        onMouseDown={this.getCaretPosition}
                        onMouseUp={this.getCaretPosition}
                        onFocus={this.handleFocus}
                        style={{
                            width: this.props.width, 
                            height: this.props.height, 
                            paddingLeft: '10px',
                            overflow: 'auto'
                        }}
                    >
                    </div>
                </div>
                <div className="info_block" style={{height: '5vh', position: 'absolute', bottom: 0, right: 16}}>
                    <span style={{float: 'right'}}>Line: {this.state.currentPositionCursor.line} | Col: {this.state.currentPositionCursor.col}</span>
                </div>
            </div>
        )
    }
}