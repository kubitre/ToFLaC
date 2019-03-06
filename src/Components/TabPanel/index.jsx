import React from 'react';

import './style.scss';

export default function TabPanel(props) {
    return (
        <div className="tab_panel_component">
            {[{
                "name":"test1",
                "active": true
        }, {
            "name":"test2",
            "active": false
        
        }].map((value, index)=>{
                return(
                    <div className={value.active ? "element_tab _active": "element_tab"} key={index}>
                        <div className="text_name">{value.name}</div>
                        <span className="closing_file"></span>
                    </div>
                )
            })}
        </div>
    )
}