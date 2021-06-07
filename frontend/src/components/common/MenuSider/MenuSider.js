import React from 'react';
import './MenuSider.css'
import { Layout, Menu, Icon } from 'antd';
import { Link } from 'react-router-dom';

const {
    Sider,
} = Layout;

const MenuSider = (props) => {
    const getMenu = () => {
        const MenuBar = (<Menu theme="dark" mode="inline" defaultSelectedKeys={props.selectedItem}>
            {props.menuModels.map(item => (
                <Menu.Item key={item.key} onClick={props.choose}>
                    <Icon type={item.iconType} />
                    <span>{item.displayText}</span>
                    <Link to={item.path} />
                </Menu.Item>
            ))}</Menu>);
        return (
            <> {MenuBar}</>
        )
    }
    return (
        <Sider trigger={null} collapsible collapsed={props.collapsed}>
            <div className="iconCollapseDiv">
                <Icon
                    type={props.collapsed ? 'menu-unfold' : 'menu-fold'}
                    onClick={props.toggle}
                />
            </div>
            {getMenu()}
        </Sider>
    );

}

export default MenuSider;