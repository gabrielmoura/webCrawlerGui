import {AddHotsTxt, AddToQueue, DeleteQueue, GetAllQueue, Start, Stop} from "../../wailsjs/go/services/CrawlerService";

export interface Url {
    url: string;
    depth: number;
}

export type URL = Url

export class QueueService {
    static async getAllQueue(): Promise<Url[]> {
        const res = await GetAllQueue()
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

    static async toggleCrawling(status: boolean): Promise<string> {
        if (!status) {
            const res = await Start()
            if (res.success) {
                return Promise.resolve(res.msg)
            }
            return Promise.reject(res.msg)
        } else {
            const res = await Stop()
            if (res.success) {
                return Promise.resolve(res.msg)
            }
            return Promise.reject(res.msg)
        }
    }

    /**
     * Get Hosts file entries and add to queue
     * @param url
     */
    static async addHotsTxt(url: string): Promise<string> {
        const res = await AddHotsTxt(url)
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject("Not implemented " + url)
    }

}