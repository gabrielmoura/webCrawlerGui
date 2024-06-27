import {ExportData, ExportQueue, ImportData, ImportQueue} from "../../wailsjs/go/services/SystemService";

export class BackupService {
    static async Export(): Promise<string> {
        const res = await ExportData()
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    static async Import(): Promise<string> {
        const res = await ImportData()
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    static async ExportQueue(): Promise<string> {
        const res = await ExportQueue()
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    static async ImportQueue(): Promise<string> {
        const res = await ImportQueue()
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }
}