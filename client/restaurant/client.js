const axios = require('axios')
const R = require('../wrappers/wrapper.js');
const { baseUrl, API_NAMES } = require('../constants/index')

class Restaurant {

    constructor() {
        this.baseWrapper = R(axios, baseUrl)

        this.crudWrapper = this.baseWrapper.crud(API_NAMES.RESTAURANT);

    }

    listMenu() {
        return this.crudWrapper.load()
    }

    createOrder(order) {
        return this.crudWrapper.create(order) 
    }

}

module.exports = new Restaurant();
