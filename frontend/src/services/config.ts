import {types} from "../../wailsjs/go/models.ts";
import {SaveWindowPosition, SaveWindowSize, SetPreferences} from "../../wailsjs/go/services/ConfigService";
import {EventsEmit} from "../../wailsjs/runtime";

export class ConfigService{
    static SavePreferences(b: types.PreferencesBehavior) {
        SetPreferences(types.Preferences.createFrom({
            behavior: b
        })).then((resp: types.JSResp) => {
            console.log('SetPreferences', resp)
        }).catch((err: any) => {
            console.error('SetPreferences', err)
        })

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
}