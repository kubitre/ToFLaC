import React from 'react';
import ElementLeftMenu from '../ElementLeftMenu';

export default LeftPanel = () => {
    console.log(this.props);

    const {elements_menu} = this.props;

    return (
        <div className="left_panel_component">
            {elements_menu.map((element, index)=>{
                <ElementLeftMenu key={index} element={element}/>
            })}
        </div>
    )
}