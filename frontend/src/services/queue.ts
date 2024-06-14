import {AddToQueue, DeleteQueue, GetAllQueue} from "../../wailsjs/go/services/CrawlingService";

export interface Url {
    url: string;
    depth: number;
}
export class QueueService {
    static async getAllQueue(): Promise<Url[]>{
        const res = await GetAllQueue()
        console.log(res)
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    static async addToQueue(url: string) {
        const res = await AddToQueue(url)
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    static async deleteQueue(url: string) {
        const res = await DeleteQueue(url)
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }
}