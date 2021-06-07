import React, { Component } from 'react';
import './App.css';
import { Route, Switch, Redirect, BrowserRouter } from 'react-router-dom';
import LoginPage from '../pages/LoginPage/LoginPage'
import HomePage from '../pages/HomePage/HomePage'

class App extends Component {
  render() {
    return (
      <div className='App'>
        <BrowserRouter>
          <Switch>
            <Route exact path="/login" component={LoginPage} />
            <PrivateRoute path="/" component={HomePage} />
          </Switch>

        </BrowserRouter>
      </div>
    );
  }
}

const PrivateRoute = ({ component: Component, ...rest }) => (
  <Route {...rest} render={(props) => (
    (localStorage.id !== "" && localStorage.role !== "")
      ? <Component {...props} />
      : <Redirect to='/login' />
  )} />
)


export default App;
