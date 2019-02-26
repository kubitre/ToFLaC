import React, {Component} from 'react';

export default class OutputBlockComponent extends Component{
    render = () => {
        return (
            this.props.isEnable ? 
            <div className="outpublock_component" style={{display: 'flex', flexDirection: 'column', height: this.props.height, width: this.props.width, marginTop: 30}}>
                <div className="messages" style={{display: 'flex', flexDirection: "row"}}>
                    <span style={{color: "orange", fontSize: '1.5em', marginLeft: '10px'}}>15 Warnings</span>
                </div>
                <div className="outer" style={{margin: '5px'}}>
                    <p>
                       1: Ошибка компиляции 
                    </p>
                </div>
            </div>
            :
            null
        )
    }
}