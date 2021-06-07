import React, { Component } from 'react';
import '../PopUpBuilder.css'
import Input from '../../UI//Input/Input'
import commonfunction from '../../../../utils/commonfucntion/commonfunction'
import Button from '../../UI/Button/Button'

class FormInsert extends Component {
    state = {
        Form: this.props.feilds,
        formIsValid: this.props.valid,
    }


    onSubmit = (event) => {
        event.preventDefault();
        const formData = {};
        for (let formElementIdentifier in this.state.Form) {
            formData[formElementIdentifier] = this.state.Form[formElementIdentifier].value;
        }

        this.props.submit(formData);
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
        console.log("formIsValid", formIsValid)
        this.setState({ Form: updatedForm, formIsValid: formIsValid });
    }


    getForm = () => {

        const formElementsArray = [];
        for (let key in this.state.Form) {
            formElementsArray.push({
                id: key,
                config: this.state.Form[key]
            });
        }

        let form = (
            <form id="forminsert">
                {formElementsArray.map(formElement => (
                    <>
                        <p>{formElement.id}:</p>
                        <div className="inputelement"><Input
                            key={formElement.id}
                            element={formElement.config}
                            changed={(event) => this.inputChangedHandler(event, formElement.id)}
                        /></div>
                    </>
                ))}
            </form>

        );
        return (<>
            {form}
            <div className="customfooter">
                <Button disabled={!this.state.formIsValid} text={this.props.btnActionName} clicked={this.onSubmit} />
                <Button text="Close" clicked={this.props.close} />
            </div>
        </>
        )
    }

    render() {
        return (
            <div>
                {this.getForm()}
            </div>
        );

    }
}
export default FormInsert;