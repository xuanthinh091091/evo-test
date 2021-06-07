import React, { Component } from 'react';
import './Table.css'
import Button from '../../UI/Button/Button'

class Table extends Component {

    getdata = (itemdata) => {
        const listItems = this.props.columns.map(
            (colunmName) =>
                <td>{itemdata[colunmName.dataIndex]}</td>
        );
        return <>
            {listItems}
            {this.props.actions && this.props.actions.length ? < td >
                <Button btnType="Order" text="Order"
                    clicked={(event) => this.onOrder(event, itemdata)} />
            </td> : ""
            }
        </>;
    }

    onOrder = (event, itemdata) => {
        event.preventDefault();
        this.props.order(itemdata);
    }

    getRows = () => {
        console.log(" props.dataSource", this.props.dataSource)
        const TableData = this.props.dataSource.map(
            (itemdata) =>
                <tr onDoubleClick={(event) => this.onView(event, itemdata)}>
                    {this.getdata(itemdata)}
                </tr>

        );
        return <> {TableData} </>;
    }

    onView = (event, itemdata) => {
        event.preventDefault();
        this.props.viewDetails(itemdata)
    }

    render() {
        const TableHeader = this.props.columns.map(
            (item) =>
                <th style={{ width: item.width ? item.width : "200px" }}>
                    {item.title}
                </th>
        );

        return (
            <div className="tableDiv">
                <table >
                    {TableHeader}
                    {this.props.actions && this.props.actions.length > 0 && <th className="action"></th>}
                    {this.getRows()}
                </table>
                {this.props.hideHelper ? "" : <div className="helptext">Double click on row to view more details</div>}
            </div>
        );
    }
}
export default Table;