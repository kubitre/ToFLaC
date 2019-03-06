import React from 'react';
import TabPanel from '../TabPanel';
import InputBlock from '../InputBlock';
import OutputBlock from '../OutputBlock';
import EmptyField from '../EmptyField';
import StatusBar from '../StatusBar';

import './style.scss';

function render_block(acrive_tab, props){
    switch(acrive_tab){
        case 0:
            return <InputBlock 
                            lines={props.caret.lines} 
                            textEditor={props.textEditor.html} 
                            textEditorActions={props.textEditorActions} 
                            caretActions={props.caretActions}/>
        case 1:
            return <OutputBlock />
        default: 
            return <EmptyField />
    }
}

export default function MainBlock(props){
    const {caret} = props;
    const active_tab = 0;
    return(
        <div className="main_block_component">
            <TabPanel />
            {render_block(active_tab, props)}
            <StatusBar caret={caret} status={{warnsAmount: 0, errorsAmount: 0}} language="golang"/>
        </div>
    )
}