import {StateCreator} from "zustand";
import {types} from "../../../wailsjs/go/models";
import {ConfigService} from "../../services/config.ts";


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
            ConfigService.SavePos(x, y)

            set({Behavior: b}, false, "window_pos")
        },
        setAsideWidth: (aside_width: number) => {
            let b = get().Behavior
            if (!b) {
                return
            }
            b.asideWidth = aside_width
            ConfigService.SavePreferences(b)
            set({Behavior: b}, false, "aside_width")
        },
        setWindow(width: number, height: number) {
            let b = get().Behavior
            if (!b) {
                return
            }
            ConfigService.SetSize(width, height, b.windowMaximised)
            set({Behavior: b}, false, "window_size")
        },
        setWindowTheme(dark: boolean) {
            let b = get().Behavior
            if (!b) {
                return
            }
            b.darkMode = dark
            ConfigService.SavePreferences(b)
            set({Behavior: b}, false, "window_theme")
        },
        updateBehavior(Behavior?: types.PreferencesBehavior) {
            set({Behavior: Behavior}, false, "update")
        }
    });
};