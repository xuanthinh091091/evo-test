const menuModels = {
  adminModels: [
    {
      key: "1",
      displayText: "All Books",
      path: "/books",
      component: "books",
      iconType: "table"
    },
    {
      key: "2",
      displayText: "Orders Details",
      path: "/orderdetails",
      component: "orders",
      iconType: "user"
    },
  ],
  userModels: [
    {
      key: "1",
      displayText: "Available Books",
      path: "/books",
      component: "books",
      iconType: "table"
    },
  ]
}

export default menuModels;
