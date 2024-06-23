import {ExportData, ImportData} from "../../wailsjs/go/services/SystemService";

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
}