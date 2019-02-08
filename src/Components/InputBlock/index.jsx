import React, {Component} from 'react';
import {Button} from 'reactstrap';
import KernelSemanticAnalys from '../KernelSemanticAnalys';

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
            }
        }
        this.handleAddNeLine = this.handleAddNeLine.bind(this);
        this.handleBackspace = this.handleBackspace.bind(this);
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

        if(
            (keyCode >= 32 && keyCode <= 100)
            ||
            (keyCode >= 1000 && keyCode <= 1200))
            {
                this.setState({
                    "text": +event.target.value,
                    "currentPositionCursor": {
                        "col": this.state.currentPositionCursor.col + 1,
                        "line": this.state.currentPositionCursor.line
                    }
                })
            }
        else if(keyCode === 8) {
            this.setState({
                "currentPositionCursor": {
                    "col": this.state.currentPositionCursor.col - 1 >= 1 ? this.state.currentPositionCursor.col - 1 : 1,
                    "line": this.state.currentPositionCursor.line
                }
            })
        }
        
    }

    handleScroll = (event) => {
        document.getElementById("#line_panel").scrollTo(document.getElementById("#line_panel").scrollHeight, document.getElementById("#line_panel").scrollHeight);
    }

    render = () => {
        return(
            <div className="inputblock_component" style={{display: 'flex', flexDirection: 'column', height: this.props.height, width: this.props.width, marginTop: 30}}>
                <div className="control_panel" style={{display: 'flex', flexDirection: 'row', height: window.innerHeight < 500 ? '15vh' : '5vh', marginLeft: 10, width: "100%"}}>
                    <Button color="success" style={{marginRight: '5px'}}>Execute</Button>
                    <Button color="danger">Stop</Button>
                </div>
                <div className="input_panel" style={{width: '100%', display: 'flex', flexDirection: 'row', height: this.props.height, paddingTop: '10px'}}>
                    <textarea id="#line_panel" 
                        readOnly 
                        name="line_numbers" 
                        value={this.state.lines} 
                        style={{width: '40px', overflow: 'hidden', textAlign: 'center'}} 
                        onChange={this.handleChangeTextEdit}/>

                    <textarea type="text" 
                                onKeyPress={this.handleAddNeLine} 
                                onScroll={this.handleScroll} 
                                onKeyDown={this.handleBackspace} 
                                onKeyUp={this.handleChangeTextEdit}
                                className="inputblock" 
                                placeholder="input here your code" 
                                style={{
                                    width: this.props.width, 
                                    height: this.props.height, 
                                    paddingLeft: '10px'}}
                                    />
                </div>
                <div className="info_block" style={{height: '5vh'}}>
                    <span style={{float: 'right'}}>Line: {this.state.currentPositionCursor.line} | Col: {this.state.currentPositionCursor.col}</span>
                </div>
            </div>
        )
    }
}