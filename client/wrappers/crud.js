module.exports = class CrudWrapper {
    constructor(httpClient, baseUrl, apiName) {
        this.httpClient = httpClient
        this.baseUrl = baseUrl
        this.apiName = apiName
    }

    async load() {
        const url = `${this.baseUrl.baseUrl}/${this.apiName}`;

        try {
            const responses = await this.httpClient.get(url);

            return responses.data;
        } catch (error) {
            console.error(`Error occured while loading ${this.apiName} at ${url}`);

            throw error;
        }
    }

    async create({ entities }) {
        const url = `${this.baseUrl.baseUrl}/${this.apiName}`;

        try {
            const responses = await Promise.all(entities.map(_ => this.httpClient.post(url, _)));

            return responses.map(_ => _.data);
        } catch (error) {
            console.error(`Error occured while creating ${this.apiName} at ${url}`);

            throw error;
        }
    }

    async createSingle({ entity }) {
        const url = `${this.baseUrl.baseUrl}/${this.apiName}`;

        try {
            const response = await this.httpClient.post(url, entity);

            return response && response.data;
        } catch (error) {
            console.error(`Error occured while creating single entity ${this.apiName} at ${url}`);

            throw error;
        }
    }

}


