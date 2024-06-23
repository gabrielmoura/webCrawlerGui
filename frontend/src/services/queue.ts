import {
    AddHotsTxt,
    AddToQueue,
    DeleteQueue,
    GetAllQueue,
    RemoveFromQueueByHost,
    Start,
    Stop
} from "../../wailsjs/go/services/CrawlerService";

export interface Url {
    url: string;
    depth: number;
}

export type URL = Url

export class QueueService {
    /**
     * Get all queue
     */
    static async getAllQueue(): Promise<Url[]> {
        const res = await GetAllQueue()
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    /**
     * Add to queue
     * @param url
     */
    static async addToQueue(url: string): Promise<string> {
        const res = await AddToQueue(url)
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    /**
     * Delete a URL from queue
     * @param url
     */
    static async deleteQueue(url: string): Promise<string> {
        const res = await DeleteQueue(url)
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    /**
     * Start or stop crawling
     * @param status boolean
     * @returns Promise<string>
     */
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

    /**
     * Remove all from queue by host
     * Attention: This method can remove all data
     * @param url
     */
    static async removeFromQueue(url: string): Promise<string> {
        const res = await RemoveFromQueueByHost(url)
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

}