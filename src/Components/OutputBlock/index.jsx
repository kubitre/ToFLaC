import React, {Component} from 'react';

import {connect} from 'react-redux';

import * as OutputActions from '../../Store/Parts/Output/actions';

import './style.scss';
import { bindActionCreators } from 'redux';

class OutputBlock  extends Component{
    constructor(props){
        super(props);

        this.state = {
            "output_tab": 2,
        }

        this.handleChangeOutput = this.handleChangeOutput.bind(this);
        this.getTypeByNum = this.getTypeByNum.bind(this);
        this.getWarningPart = this.getWarningPart.bind(this);
    }
    handleChangeOutput = (event) => {
        this.props.outAct.changeOutputType(event.target.getAttribute("ident"))
    }

    getTypeByNum = (num) => {
        switch (num) {
            case -2:
                return "Error"
            case 0:
                return "NONTYPE"
            case 2:
                return "INT"
            case 4:
                return "FLOAT"
            case 6:
                return "IDENTIFIER"
            case 8:
                return "ENDSTATEMENT"
            case 10:
                return "POINTER"
            case 12:
                return "SPACE"
            case 14:
                return "NEW LINE"
            case 16:
                return "COMMA"
            default:
                return ""
        }
    }

    getStringViewByTypeToken(token){
        switch(token.Type){
            case 2:
                return "int";
            case 4:
                return "float";
            case 6:
                return token.Value;
            case 8:
                return ";"
            case 10:
                return "*";
            case 12:
                return " ";
            case 14:
                return "\n";
            case 16:
                return ",";
            default:
                return "";
        }
    }

    getRepairThing(tokens){
        let decryptString = "";
        for(let i = 0; i < tokens.length; i ++){
            decryptString += this.getStringViewByTypeToken(tokens[i])
        }
        return decryptString;
    }

    getWarningPart(start, end, type){
        // console.log(this.props.storeInput)
        let value = this.props.storeInput.code.substring(start, end)
        let tpe = ""
        if (type === -4) {
            tpe = "NONTYPE "
        }
        return tpe + value
    }

    renderOutputBody(type){
        const {LexerAnalysPart, SyntaxAnalysPart} = this.props.store.output;

        switch(type){
            case 1:
                return (
                    <div className="tokens">
                        {LexerAnalysPart.Tokens != null ? 
                            LexerAnalysPart.Tokens.map((token, index) => {
                                return(
                                    <div className="token_container">
                                        <div className="token" key={index}>{this.getTypeByNum(token.Type)} {token.Value!=null ? "|" + token.Value: null}</div>
                                        {index - 1 === LexerAnalysPart.Tokens.length ?
                                        null
                                            :
                                            <div className="arrow">-></div>
                                        }
                                    </div>
                                )
                            })
                            :
                            null
                        }
                    </div>
                )
            case 2:
                return (
                    <div className="warnings">
                        {LexerAnalysPart.Errors != null ?
                            LexerAnalysPart.Errors.map((error, index) => {
                                return (
                                    <div className="warning" key={index}>At Line:{error.Token.Line}, Column: {error.Token.Column} handle error of Lexer Analys: {error.Message}</div>
                                )
                            })
                            :
                            null
                        }
                    </div>
                )
            case 3:
                return(
                    <div className="errors">
                        {SyntaxAnalysPart.Errors != null ?
                            SyntaxAnalysPart.Errors.map((error, index) => {
                                return(
                                    <div className="error" key={index}>At Line: {error.Token.Line} handle error of Syntax Analys: {error.Message}</div>
                                )
                            }):
                            null
                        }
                    </div>
                )

            case 4:
                return (
                    <div className="repair_part">
                        {
                            SyntaxAnalysPart.repair !== null ? 
                                <div className="repair_part">Repair: {this.getRepairThing(SyntaxAnalysPart.repair)}</div>
                                :
                                null
                        }
                    </div>
                )
            default:
                return (
                    <div className="not_output" style={{color: 'white'}}>You should choosed a type of output</div>
                )
        }
    }

    render = () => {
        // console.log(this.state.output_tab);
        // console.log(this.props.store.output_data);
        // document.getElementsByClassName("field_for_out").innerHTML = text_with_analys;
        console.log(this.props.store)
        const {type} = this.props.store;
        
        return (
            <div className="output_block_component">
                <div className="header">
                    <div className="result_name">Результат </div>    
                </div>
                <div className="output_separator">
                    <div className="button_tokens" ident={1} onClick={this.handleChangeOutput}>Токены</div>
                    <div className="button_warnings" ident={2} onClick={this.handleChangeOutput}>Warnings</div>
                    <div className="button_erros" ident={3} onClick={this.handleChangeOutput}>Errors</div>
                    <div className="button_repair" ident={4} onClick={this.handleChangeOutput}>Repair</div>
                </div>
                <div className="field_for_out" style={{padding: `15px`}}>
                    {this.renderOutputBody(type)}
                </div>
            </div>
        )
    }
}



function mapStateToProps (state) {
    return {
        // store: state.outputBlockState,
        // storeInput: state.inputBlockState
        store: state.OSI_outputState
    }
}

function mapDispatchToProps(dispatch) {
    return {
        outAct: bindActionCreators(OutputActions, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(OutputBlock)