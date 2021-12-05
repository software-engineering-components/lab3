const CrudWrapper = require('./crud.js');

module.exports = (httpClient, baseUrl) => {
    return {
        crud: (apiName) => {
            const wrapper = new CrudWrapper(httpClient, baseUrl, apiName);

            return wrapper;
        },
    };
}
