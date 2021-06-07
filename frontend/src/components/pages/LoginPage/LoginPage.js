import React, { Component } from 'react';
import FormBuilder from '../../common/FormBuilder/FormBuilder'
import LoginFormModel from '../../../models/LoginFormModel'
import axios from 'axios';

class LoginForm extends Component {

    login = (formData) => {
        axios.post('/login', formData)
            .then(response => {
                console.log(response)
                localStorage.setItem("id", response.data.id)
                localStorage.setItem("username", formData.username)
                localStorage.setItem("role", response.data.role)
                this.props.history.push('/')
            })
            .catch(error => {
                console.log("error", error)
            });
    }

    render() {
        return (
            <div>
                <FormBuilder submitForm={this.login} FormModel={LoginFormModel} buttonText="Sign in" />
            </div>
        );
    }
}
export default LoginForm;