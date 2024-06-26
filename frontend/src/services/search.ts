import {GetStatistics, SearchWords} from "../../wailsjs/go/services/CrawlerService";

export interface MetaData {
    og: { [key: string]: string };
    keywords: string[];
    manifest: string;
    ld: string;
}

export interface Page {
    url: string;
    links: string[];
    title: string;
    description: string;
    meta?: MetaData;
    visited: boolean;
    timestamp: Date;
    words: { [key: string]: number };
}

export interface Statistics {
    total_pages: number
    total_pages_per_host?: Record<string, number>
}

export class SearchService {
    static async searchWords(query: string): Promise<Array<Page>> {
        // separe as palavras por espaço
        const words = query.split(' ')

        const rest = await SearchWords(words)
        if (rest.success) {
            return Promise.resolve(rest.data)
        }
        return Promise.reject(rest.msg)
    }

    static async GetStatistics(): Promise<Statistics> {
        const rest = await GetStatistics()
        if (rest.success) {
            return Promise.resolve(rest.data)
        }
        return Promise.reject(rest.msg)
    }
}