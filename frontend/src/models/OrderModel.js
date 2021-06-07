const orderModel = {
    quantity: {
        elementType: 'input',
        elementConfig: {
            type: 'text',
            placeholder: 'Quantity'
        },
        value: '',
        validation: {
            required: true,
            isNumeric: true,
        },
        valid: false,
        touched: false,
        errMessage: 'Quantity must be a number',
    },
    shipping_details: {
        elementType: 'input',
        elementConfig: {
            type: 'text',
            placeholder: 'Shipping Details'
        },
        value: '',
        validation: {
            required: true,
        },
        valid: false,
        touched: false,
        errMessage: 'Shipping Details cannot be empty',
    }
}
export default orderModel