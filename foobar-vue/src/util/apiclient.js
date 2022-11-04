import axios from "axios";

export default class ApiClient {
    basePath;
    constructor(basePath) {
        this.basePath = basePath
        this.axios = axios.create({
            baseURL:basePath,
            headers: {
                'Content-Type': 'application/json'
            }
        })
    }

    async post(url, payload){
        return this.axios.post(url, payload)
    }

    async get(url, payload){
        return this.axios.get(url, payload)
    }

    async getWithParam(url, data){
        return this.axios.get(url, { params : data})
    }

    async delete(url, payload){
        return this.axios.delete(url, payload)
    }

}