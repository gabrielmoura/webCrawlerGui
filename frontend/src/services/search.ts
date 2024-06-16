import {Search, SearchWords} from "../../wailsjs/go/services/CrawlerService";


export class SearchService {
    /**
     * Pesquisa páginas por título, descrição ou conteúdo
     * @param query - Termo de pesquisa
     */
    static async search(query: string) {
        const rest = await Search(query)
        if (rest.success) {
            return Promise.resolve(rest.data)
        }
        return Promise.reject(rest.msg)
    }

    static async searchWords(query: string) {
        // separe as palavras por espaço
        const words = query.split(' ')

        const rest = await SearchWords(words)
        if (rest.success) {
            return Promise.resolve(rest.data)
        }
        return Promise.reject(rest.msg)
    }
}