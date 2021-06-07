import React, { Component } from 'react';
import './Books.css';
import Table from '../../common/UI/Table/Table'
import PopUpBuilder from '../../common/PopUpBuilder/PopUpBuilder'
import axios from 'axios';
import OrderModel from '../../../models/OrderModel'

export const columns = [{
  title: 'Name',
  dataIndex: 'name',
  key: 'name',
  width: "350px",
}, {
  title: 'Author',
  dataIndex: 'author',
  key: 'author',
  width: "300px",
}, {
  title: 'Publisher',
  dataIndex: 'publisher',
  key: 'publisher',
  width: "300px",
}, {
  title: 'Price',
  dataIndex: 'price',
  key: 'price',
  width: "200px",
}, {
  title: 'Stock',
  dataIndex: 'stock',
  key: 'stock',
  width: "200px",
}
];

class Books extends Component {
  state = {
    books: [],
    openModal: false,
    orderItem: {},
    modalData: {}
  };

  order = (item) => {
    console.log(item)
    var modaldata = {
      type: "insert",
      title: "New Order For Book: " + item.name,
      item: {}
    }
    this.setState({ openModal: true, modalData: modaldata, orderItem: item })
  }


  viewDetails = (item) => {
    var modaldata = {
      type: "view",
      title: "Book details",
      item: {
        "Name": item.name,
        "Author": item.author,
        "Publisher": item.publisher,
        "Total Page": item.total_page,
        "Published Date": item.published_date,
        "Price": item.price,
        "Stock": item.stock,
      },
    }
    this.setState({ openModal: true, modalData: modaldata })
  }

  closeModal = () => {
    this.setState({ openModal: false })
  }

  componentDidMount() {
    var url = '/user/booksavailableforsale'
    if (localStorage.role === "ADMIN") {
      url = '/admin/books'
    }
    axios.get(url, { 'headers': { 'userID': localStorage.id } })
      .then(response => {
        console.log(response)
        this.setState({ books: response.data })
      })
      .catch(error => {
        console.log("error", error)
      });
  }

  Submit = (item) => {
    if (item.quantity > this.state.orderItem.stock) {
      alert("quantity cannot be greater than " + this.state.orderItem.stock)
      return
    }
    const data = {
      users_id: parseInt(localStorage.id),
      books_id: this.state.orderItem.id,
      quantity: parseInt(item.quantity),
      shipping_details: item.shipping_details,
    }
    axios.post('/user/book/order', data, { 'headers': { 'userID': localStorage.id } })
      .then(response => {
        console.log(response)
        this.closeModal()
      })
      .catch(error => {
        console.log("error", error)
        alert("cannot make order " + error)
      });
  }

  render() {
    return (
      <div className="Userdiv">
        <div className="pageTitle">Books</div>
        <div className="buttondiv">
          <Table columns={columns} dataSource={this.state.books}
            viewDetails={this.viewDetails}
            order={this.order}
            actions={localStorage.role === "USER" ? ["Order"] : []} />
        </div>
        <PopUpBuilder delete={this.deleteBook} submit={this.Submit} data={this.state.modalData}
          closeModal={this.closeModal} visible={this.state.openModal} feilds={OrderModel} />
      </div>

    );
  }
}

export default Books;