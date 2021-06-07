const LoginFormModel = {
    username: {
        elementType: 'input',
        elementConfig: {
            type: 'text',
            placeholder: 'Username'
        },
        value: '',
        validation: {
            required: true,
            minLength: 4,
        },
        valid: false,
        touched: false,
        errMessage: 'The username must be at least 4 characters in length',
    },
    password: {
        elementType: 'input',
        elementConfig: {
            type: 'password',
            placeholder: 'Password'
        },
        value: '',
        validation: {
            required: true,
            minLength: 6,
        },
        valid: false,
        touched: false,
        errMessage: 'Password must have at least 6 characters in length',
    }
}
export default LoginFormModel