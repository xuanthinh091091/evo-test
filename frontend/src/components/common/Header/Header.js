import React from 'react';
import './Header.css'
import {Icon} from 'antd';

const Header = (props) => {
    return (
        <div className="Header">
            <div className="usernameDiv">
                <div className="iconLogoutDiv" onClick={props.logout}>  
                    <Icon type="logout" />
                </div>
                <div className="username">{localStorage.username}</div>
            </div>
        </div>
    );
}

export default Header;