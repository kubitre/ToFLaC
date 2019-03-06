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
                        backgroundSize: 'cover',
                        backgroundRepeat: 'no-repeat',
                        width: '35px',
                        height: '35px'
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