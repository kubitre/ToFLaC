import React, {Component} from 'react';
import Navbar from '../../Components/Navbar';
import InputBlockComponent from '../../Components/InputBlock';
import OutputBlockComponent from '../../Components/OutputBlock';
import { connect } from 'react-redux';


class MainPageComponent extends Component{
    render = () => {
        const {enabledOutputBlock} = false
        return (
            <div className="main_page_container" style={{display: 'flex', flexDirection: 'column'}}>
                {window.innerHeight < 500 ? 
                    <Navbar height="10vh"/>
                    :
                    <Navbar height="20vh"/>
                }
                <div className="blocks_row" style={{display: 'flex', flexDirection: "column", marginLeft: '10px'}}>
                    <InputBlockComponent height="60vh" width="97vw"/>
                    <OutputBlockComponent height="40vh" width="90vw" isEnabled={enabledOutputBlock}/>
                </div>
            </div>
        )
    }
}

function mapStateToProps (state) {
    return {
        user: state.user
    }
}

export default connect(mapStateToProps)(MainPageComponent)