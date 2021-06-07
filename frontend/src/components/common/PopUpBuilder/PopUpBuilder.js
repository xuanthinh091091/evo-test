import React, { Component } from 'react';
import { Modal } from 'antd';
import Button from '../UI/Button/Button'
import './PopUpBuilder.css'
import FormInsert from './FormInsert/FormInsert'

class PopUpBuilder extends Component {

    details = () => {
        var ElementsArray = []
        for (var item in this.props.data.item) {
            ElementsArray.push({
                key: item,
                value: this.props.data.item[item]
            });
        }

        let items = (<>
            {ElementsArray.map(element => (
                <div className="textdiv">
                    <div className="text">{element.key}:</div>
                    <div className="value">{element.value}</div>
                </div>
            ))}
        </>
        );

        return (
            <>
                {items}
                <div className="customfooter">
                    <Button text="Close" clicked={this.props.closeModal} />
                </div>
            </>
        )
    }

    getForm = () => {
        let newfi = this.props.feilds
        for (let item in newfi) {
            if (this.props.data.item[item]) {
                newfi[item].value = this.props.data.item[item]
                newfi[item].valid = true
            }
            else {
                newfi[item].value = ""
                newfi[item].valid = false
            }
        }
        return (
            <> <FormInsert feilds={newfi} valid={false} btnActionName={"Submit"} close={this.props.closeModal} submit={this.props.submit} /> </>
        )
    }

    getbody = () => {
        let inputElement = null;
        switch (this.props.data.type) {
            case 'insert':
                inputElement = this.getForm(true)
                break;
            default:
                inputElement = this.details()
        }
        return (<>{inputElement}</>);

    }

    render() {
        return (
            <div>
                <Modal
                    title={this.props.data.title}
                    visible={this.props.visible}
                    onCancel={this.props.closeModal}
                    footer={null}
                    destroyOnClose={true}
                >
                    {this.getbody()}
                </Modal>
            </div>
        );

    }
}
export default PopUpBuilder;