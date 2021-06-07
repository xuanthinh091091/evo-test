import React, { Component } from 'react';
//import Input from '../../../components/UI/Input/Input'
import Input from '../UI/Input/Input'
import Button from '../UI/Button/Button'
import './FormBuilder.css'
import commonfunction from '../../../utils/commonfucntion/commonfunction'

class FormBuilder extends Component {
    state = {
        Form: this.props.FormModel,
        formIsValid: false,
        loading: false
    }

    inputChangedHandler = (event, inputIdentifier) => {
        const updatedForm = {
            ...this.state.Form
        };
        const updatedFormElement = {
            ...updatedForm[inputIdentifier]
        };
        updatedFormElement.value = event.target.value;
        updatedFormElement.valid = commonfunction.checkValidity(updatedFormElement.value, updatedFormElement.validation);
        updatedFormElement.touched = true;
        updatedForm[inputIdentifier] = updatedFormElement;

        let formIsValid = true;
        for (let inputIdentifier in updatedForm) {
            formIsValid = updatedForm[inputIdentifier].valid && formIsValid;
        }
        this.setState({ Form: updatedForm, formIsValid: formIsValid });
    }

    onSubmitForm = (event) => {
        event.preventDefault();
        const formData = {};
        for (let formElementIdentifier in this.state.Form) {
            formData[formElementIdentifier] = this.state.Form[formElementIdentifier].value;
        }
        this.props.submitForm(formData);
    }

    render() {
        const formElementsArray = [];
        for (let key in this.state.Form) {
            formElementsArray.push({
                id: key,
                config: this.state.Form[key]
            });
        }
        let form = (
            <form onSubmit={this.onSubmitForm}>
                {formElementsArray.map(formElement => (
                    <Input
                        key={formElement.id}
                        element={formElement.config}
                        changed={(event) => this.inputChangedHandler(event, formElement.id)}
                    />
                ))}
                <Button btnType="SubmitForm" disabled={!this.state.formIsValid} text={this.props.buttonText} />
            </form>

        );

        return (
            <div className='Form_inner'>
                {form}
            </div>
        );
    }
}

export default FormBuilder;