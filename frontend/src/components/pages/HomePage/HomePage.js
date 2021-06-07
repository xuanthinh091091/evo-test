import React, { Component } from 'react';
import './HomePage.css'
import 'antd/dist/antd.css'
import { Layout } from 'antd';
import Header from '../../common/Header/Header'
import Books from '../Books/Books'
import Orders from '../Orders/Orders'
import { Route, Redirect } from 'react-router-dom';
import MenuSider from '../../common/MenuSider/MenuSider'
import menuModels from '../../../models/MenuModels'

const {
  Content
} = Layout;
class HomePage extends Component {
  state = {
    collapsed: false,
  };

  getDefaultComponent = () => {
    var itemkey = localStorage.adminchooseItem;
    if (!itemkey) {
      localStorage.setItem("adminchooseItem", "1")
    }
    var models = menuModels.userModels
    if (localStorage.role === "ADMIN") {
      models = menuModels.adminModels
    }
    console.log(menuModels)
    //console.log(parseInt(localStorage.adminchooseItem) - 1)
    //console.log()
    var path = models[parseInt(localStorage.adminchooseItem) - 1].path;
    if (this.props.location.pathname === "/")
      return <Redirect from="/" to={path} />
    else
      this.props.history.push("/pagenotfound")
  }

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  }
  logout = () => {
    localStorage.removeItem("id")
    localStorage.removeItem("role")
    localStorage.removeItem("adminchooseItem")
    this.props.history.push('/login');
  }

  choose = (item) => {
    localStorage.setItem("adminchooseItem", item.key)
  }

  render() {

    return (<>
      <Layout>
        <MenuSider selectedItem={localStorage.adminchooseItem} choose={this.choose} toggle={this.toggle} collapsed={this.state.collapsed} menuModels={localStorage.role === "ADMIN" ? menuModels.adminModels : menuModels.userModels} />
        <Layout>
          <Header logout={this.logout} />
          <Content style={{ margin: '24px 16px', padding: 24, background: '#fff', minHeight: 580, }}>
            <Route exact path="/books" component={Books} />
            <Route exact path="/orderdetails" component={Orders} />
            <Route exact path="/" component={this.getDefaultComponent} />
          </Content>
        </Layout>
      </Layout>
    </>
    );
  }

}



export default HomePage;