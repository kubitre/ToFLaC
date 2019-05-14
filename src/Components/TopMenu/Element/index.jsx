import React, {Component} from 'react';

import {connect} from 'react-redux';

import * as ModalActions from '../../../Store/Parts/Modal/actions';
import * as FileActions from '../../../Store/Actions/file_try';

import './style.scss';
import { bindActionCreators } from 'redux';

class TopMenuElement extends Component {
    constructor(props){
        super(props);

        this.state = {
            "click": false,
        }

        this.workWithModalWindow = this.workWithModalWindow.bind(this);
        this.workWithNewFile = this.workWithNewFile.bind(this);
        this.workWithOpenFile = this.workWithOpenFile.bind(this);
    }

    handleClick(event){
        if (event.target.getAttribute("action") != null) {
            this.handleInlineClick(event)
            return
        }
        this.setState({click: !this.state.click})
    }

    handleInlineClick(event){
        const action = event.target.getAttribute("action")
        const params = event.target.getAttribute("params")

        switch(action){
            case "modal":
                this.workWithModalWindow(params);
                this.props.modalWindowActions.OpenModalWindow();
                return
            case "open":
                this.workWithOpenFile(params);
                return
            default:
                return
        }
    }

    workWithModalWindow(params){
        switch(params){
            case "help": 
                this.props.modalWindowActions.SetModalHelp()
                return
            case "stag":
                this.props.modalWindowActions.SetModalTask()
                return
            case "gram":
                this.props.modalWindowActions.SetModalGrammar()
                return
            default:
                return;
        }
    }

    workWithOpenFile(params){
        switch(params){
            case "example1":
                console.log("handle open example1")
                this.props.fileOpen.OpenFileExample(1)
                return
            case "example2":
                console.log("handle open example1")
                this.props.fileOpen.OpenFileExample(2)
                return
            case "example3":
                console.log("handle open example1")
                this.props.fileOpen.OpenFileExample(3)
                return
            case "example4":
                console.log("handle open example1")
                this.props.fileOpen.OpenFileExample(4)
                return
            case "example5":
                console.log("handle open example1")
                this.props.fileOpen.OpenFileExample(5)
                return
            default:
                return
        }
    }

    workWithNewFile(params){

    }

    render = () => {
        const {elementsInline, name, action, params} = this.props.elemen;
        return (
            <div className="topmenu_element" >
                <div className="head" onClick={this.handleClick.bind(this)} action={action} params={params}>{name}</div>
                {this.state.click ? 
                    <div className="sub">
                        {elementsInline != null ? 
                            elementsInline.map((element, index) => {
                                return (
                                    <div className="content" onClick={this.handleInlineClick.bind(this)} action={element.action} params={element.params}>{element.name}</div>
                                )
                            })
                            :
                            null
                        }
                    </div>
                    :
                    null
                }
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
        // store: state.outputBlockState,
        // storeInput: state.inputBlockState

    }
}

function mapDispatchToProps(dispatch) {
    return {
        modalWindowActions: bindActionCreators(ModalActions, dispatch),
        fileOpen: bindActionCreators(FileActions, dispatch),
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(TopMenuElement)
