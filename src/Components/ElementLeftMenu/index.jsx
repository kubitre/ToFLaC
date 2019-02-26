import React from 'react';

import './style.scss';

export default ElementLeftMenu = () => {
    console.log(this.props);
    const {img, name, visible_img} = this.props.element;
    return (
        <div className="element_left_panel_menu">
            {visible_img ? 
                <div className="img_element" 
                    style={{
                        background: `url(${img})`,
                        backgroundSize: 'cover',
                        backgroundRepeat: 'no-repeat',
                        width: '75px',
                        height: 'auto'
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