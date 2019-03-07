import React, {Component} from 'react';

import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as linesActions from '../../Store/Actions/linesshower';

class LinesShower extends Component{
    render = () => {
        const lines = [];
        return (
            <div className="lines_shower_block">
                {lines.map((line, index) => {
                    return(
                        <div className="element_left_side" key={index}>{line}</div>
                    )
                })}   
            </div>
        )
    }
}


function mapStateToProps (state) {
    return {
        lines: state.linesShowerState
    }
}

function mapDispatchToProps(dispatch) {
    return {
        linesActions: bindActionCreators(linesActions, dispatch)
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(LinesShower)

