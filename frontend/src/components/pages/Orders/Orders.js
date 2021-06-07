import React, { Component } from 'react';
import './Orders.css';
import Table from '../../common/UI/Table/Table'
import axios from 'axios';

export const columns = [{
  title: 'ID',
  dataIndex: 'id',
  key: 'id',
  width: "50px",
}, {
  title: 'User Name',
  dataIndex: 'user_name',
  key: 'user_name',
  width: "150px",
}, {
  title: 'Book Name',
  dataIndex: 'book_name',
  key: 'book_name',
  width: "350px",
}, {
  title: 'Quantity',
  dataIndex: 'quantity',
  key: 'quantity',
  width: "100px",
}, {
  title: 'Shipping Details',
  dataIndex: 'shipping_details',
  key: 'shipping_details',
  width: "400px",
}, {
  title: 'Created At',
  dataIndex: 'created_at',
  key: 'created_at',
  width: "200px",
}];

class Orders extends Component {
  state = {
    books: [],
    openModal: false,
    orderItem: {},
    modalData: {}
  };

  componentDidMount() {
    var url = '/admin/orders'
    axios.get(url, { 'headers': { 'userID': localStorage.id } })
      .then(response => {
        console.log(response)
        this.setState({ books: response.data })
      })
      .catch(error => {
        console.log("error", error)
      });
  }

  render() {
    return (
      <div className="Userdiv">
        <div className="pageTitle">Orders</div>
        <div className="buttondiv">
          <Table columns={columns} dataSource={this.state.books}
            viewDetails={this.viewDetails}
            hideHelper={true} />
        </div>
      </div>

    );
  }
}

export default Orders;