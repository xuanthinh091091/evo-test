import React from 'react';

import classes from './Input.module.css';

const input = (props) => {
    let inputElement = null;

    const inputClasses = [classes.InputElement];
    const errMesClasses = [classes.ErrMess];
    if (!props.element.valid && props.element.validation && props.element.touched) {

        inputClasses.push(classes.Invalid);
        errMesClasses.push(classes.Invalid);
    }

    switch (props.element.elementType) {
        case ('input'):
            inputElement = <input className={inputClasses.join(' ')} {...props.element.elementConfig}
                onChange={props.changed} onEnded={props.validate} value={props.element.value}
            />;
            break;
        default:
            inputElement = <div></div>
    }

    return (
        <>
            {inputElement}
            <div className={errMesClasses.join(' ')}>{props.element.errMessage}</div>
        </>
    );

};

export default input;