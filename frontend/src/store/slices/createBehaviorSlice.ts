import {StateCreator} from "zustand";
import {types} from "../../../wailsjs/go/models";
import {EventsEmit} from "../../../wailsjs/runtime";
import {SaveWindowPosition, SaveWindowSize, SetPreferences} from "../../../wailsjs/go/services/ConfigService";
// import {SaveWindowPosition, SaveWindowSize, SetPreferences} from "@wails/go/services/PreferencesService";
// import {types} from "@wails/go/models";
// import {EventsEmit} from "@wails/runtime";


function SavePreferences(b: types.PreferencesBehavior) {
    SetPreferences(types.Preferences.createFrom({
        behavior: b
    })).then((resp: types.JSResp) => {
        console.log('SetPreferences', resp)
    }).catch((err: any) => {
        console.error('SetPreferences', err)
    })

}

function SavePos(x: number, y: number) {
    console.log('SavePos', x, y)
    // WindowSetPosition(x, y)
    SaveWindowPosition(x, y).then(() => {
        console.log('SaveWindowPosition')
    }).catch((err: any) => {
        console.error('SaveWindowPosition', err)
    })
}

function SetSize(width: number, height: number, maximised: boolean) {
    SaveWindowSize(width, height, maximised).then(() => {
        console.log('SaveWindowSize')
    }).catch((err: any) => {
        console.error('SaveWindowSize', err)
    })
}


interface BehaviorSlice {
    Behavior?: types.PreferencesBehavior
}

interface BehaviorSliceActions {
    setWelcomed: () => void
    setAsideWidth: (aside_width: number) => void

    setWindow(width: number, height: number): void

    setWindowPosition(x: number, y: number): void

    setWindowTheme(dark: boolean): void

    updateBehavior(Behavior: types.PreferencesBehavior
    ): void
}

export type BehaviorStore = BehaviorSlice & BehaviorSliceActions;

export const createBehaviorSlice: StateCreator<
    BehaviorSlice & BehaviorSliceActions,
    [["zustand/devtools", never]],
    []
> = (set, get) => {
    return ({
        setWelcomed: () => {
            let b = get().Behavior
            if (b?.welcomed) {
                b.welcomed = true
                set({Behavior: b}, false, "welcomed")
            }
        },
        setWindowPosition(x: number, y: number) {
            let b = get().Behavior
            if (!b) {
                return
            }
            b.windowPosX = x
            b.windowPosY = y
            // SavePreferences(b)
            SavePos(x, y)
            EventsEmit('window-move', {x: x, y: y})
            set({Behavior: b}, false, "window_pos")
        },
        setAsideWidth: (aside_width: number) => {
            let b = get().Behavior
            if (!b) {
                return
            }
            b.asideWidth = aside_width
            SavePreferences(b)
            set({Behavior: b}, false, "aside_width")
        },
        setWindow(width: number, height: number) {
            let b = get().Behavior
            if (!b) {
                return
            }
            SetSize(width, height, b.windowMaximised)
            EventsEmit('window-resize', {width: width, height: height})
            set({Behavior: b}, false, "window_size")
        },
        setWindowTheme(dark: boolean) {
            let b = get().Behavior
            if (!b) {
                return
            }
            b.darkMode = dark
            SavePreferences(b)
            set({Behavior: b}, false, "window_theme")
        },
        updateBehavior(Behavior?: types.PreferencesBehavior) {
            set({Behavior: Behavior}, false, "update")
        }
    });
};