import {types} from "../../wailsjs/go/models.ts";
import {
    AddToBlacklist,
    GetBlacklist,
    GetPreferences,
    SaveWindowPosition,
    SaveWindowSize,
    SetPreferences,
} from "../../wailsjs/go/services/ConfigService";
import {EventsEmit, Hide} from "../../wailsjs/runtime";

export class ConfigService {
    static SavePreferences(b: types.PreferencesBehavior) {
        SetPreferences(types.Preferences.createFrom({
            behavior: b
        })).then((resp: types.JSResp) => {
            console.log('SetPreferences', resp)
        }).catch((err: any) => {
            console.error('SetPreferences', err)
        })

    }

    static async SaveAllPreferences(g: types.Preferences) {
        const res = await SetPreferences(types.Preferences.createFrom(g));
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }

    static SavePos(x: number, y: number) {
        console.log('SavePos', x, y)
        // WindowSetPosition(x, y)
        SaveWindowPosition(x, y).then(() => {
            console.log('SaveWindowPosition')
            EventsEmit('window-move', {x: x, y: y})
        }).catch((err: any) => {
            console.error('SaveWindowPosition', err)
        })
    }

    static SetSize(width: number, height: number, maximised: boolean) {
        SaveWindowSize(width, height, maximised).then(() => {
            console.log('SaveWindowSize')
            EventsEmit('window-resize', {width: width, height: height})

        }).catch((err: any) => {
            console.error('SaveWindowSize', err)
        })
    }

    static async Get(): Promise<types.Preferences> {
        const res = await GetPreferences();
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    static HideWindow() {
        Hide()
    }

    static async GetBlacklist(): Promise<string[]> {
        const res = await GetBlacklist();
        if (res.success) {
            return Promise.resolve(res.data)
        }
        return Promise.reject(res.msg)
    }

    static async AddToBlacklist(url: string): Promise<string> {
        const res = await AddToBlacklist(url);
        if (res.success) {
            return Promise.resolve(res.msg)
        }
        return Promise.reject(res.msg)
    }
}