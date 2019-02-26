import React from 'react';

import './style.scss';

export default UserBlock = () =>{ 
    const {username, avatar} = this.props;
    console.log(this.props);
    let avatar_local = "./32496864.png";
    return (
        <div className="user_block_managment">
            <div className="user_avatar" 
                style={{
                    background: `url(${avatar_local})`,
                    backgroundRepeat: 'no-repeat',
                    backgroundSize: 'cover',
                    width: '70px',
                    height: 'auto'
                }}
            />
            <div className="user_name">{username}</div>
        </div>
    )
}