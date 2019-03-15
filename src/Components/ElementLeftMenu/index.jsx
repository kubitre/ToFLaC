import React from 'react';

import './style.scss';

export default function ElementLeftMenu(props) {
    // console.log(this.props);
    const {img, name, visible_img} = props.element;
    return (
        <div className="element_left_panel_menu">
            {visible_img ? 
                <div className="img_element" 
                    style={{
                        background: `url(${img})`,
                    }}
                />
                :
                <div className="name_element">
                    {name}
                </div>
            }
        </div>
    )
}