import {
    AddHotsTxt,
    AddToQueue,
    DeleteAllFailed,
    DeleteFailed,
    DeleteQueue,
    GetAllFailed,
    GetAllQueue,
    GetTreePages,
    RemoveFromQueueByHost,
    Start,
    Stop,
} from "../../wailsjs/go/services/CrawlerService";
import {TreeNode} from "../components/TreeLink.tsx";

export interface Url {
    url: string;
    depth: number;
}

export type URL = Url

export interface Pagination {
    pageNumber: number;
    pageSize: number;
}

export interface FailedType {
    url: string;
    reason?: string;
}

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
     * @param prefix
     */
    static async removeFromQueue(prefix: string): Promise<string> {
        const res = await RemoveFromQueueByHost(prefix)
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    static async getTreePages({pageNumber, pageSize = 10}: Pagination): Promise<Array<TreeNode>> {
        const res = await GetTreePages(pageNumber, pageSize)
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    static async getAllFailed(): Promise<FailedType[]> {
        const res = await GetAllFailed()
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    static async deleteFailed(url: string): Promise<string> {
        const res = await DeleteFailed(url)
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    static async deleteAllFailed(): Promise<string> {
        const res = await DeleteAllFailed("")
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }
}
