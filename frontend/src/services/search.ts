import {Search} from "../../wailsjs/go/services/CrawlingService";


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
}