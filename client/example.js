const restaurant = require('./restaurant/client')

// Scenario 1: Display available menu
restaurant.listMenu()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available menu:');

        console.log(JSON.stringify(list))
    })
    .catch((error) => {
        console.log(`Problem listing available menu: ${error.message}`);
    });

// Scenario 2: Create new order
restaurant.createOrder({
  tableNamber: 1,
  items: [
    { itemId: 1, quantity: 3 },
    { itemId: 2, quantity: 2 }
  ],
})
    .then((order) => {
        console.log('=== Scenario 2 ===');
        console.log('Create order response:');

        console.log(JSON.stringify(order))
    })
    .catch((error) => {
        console.log(`Problem creating a new order: ${error.message}`);
    });
