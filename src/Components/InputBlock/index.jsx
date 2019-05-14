import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
// import {parseTmTheme} from 'monaco-themes';

import './style.scss';
// import ContentEditable from 'react-contenteditable';
// import BlockHoverInformation from '../BlockHover';

// import * as caretActions from '../../Store/Actions/caret';
// import * as inputActions from '../../Store/Actions/inputblock';

import * as InputActions from '../../Store/Parts/Input/actions';
import * as StatusAction from '../../Store/Parts/StatusBar/actions';

import MonacoDiffEditor from 'react-monaco-editor'
var dataG = null

class InputBlock extends Component {
    constructor(props){
        super(props);
        // this.monaco = new MonacoDiffEditor()
        this.onChange= this.onChange.bind(this);
    }
    editorDidMount(editor, monaco) {
        console.log('editorDidMount', editor);
        editor.focus();
    }
    onChange(newValue, e) {
        // console.log('onChange', newValue, e);
        // this.props.setNewText.addTextToInputBlock(newValue);
        // console.log(e)
        // this.props.statusAct.AddLineCol({line: 0, column: 0})
        this.props.inpAct.SetupSource(newValue)
        // console.log("CODE:", this.props.store.code)

    }

    handleInputText = (event) => {

    }

    render = () => { 
        const {code} = this.props.store;
        const options = {
            selectOnLineNumbers: true
          };
          
        return(
            <div className="input_block_component" id="text">
                <MonacoDiffEditor
                    language="c++"
                    theme="vs-dark"
                    original={"test"}
                    value={code}
                    options={options}
                    onChange={this.onChange}
                    editorDidMount={this.editorDidMount}
                />
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
        store: state.ISI_inputState,
    }
}

function mapDispatchToProps(dispatch) {
    return {
        inpAct: bindActionCreators(InputActions, dispatch),
        statusAct: bindActionCreators(StatusAction, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(InputBlock)