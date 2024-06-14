import {StateCreator} from "zustand";
import {types} from "../../../wailsjs/go/models";


interface GeneralSlice {
    General?: types.PreferencesGeneral
}

interface GeneralSliceActions {
    updateGeneral(General: types.PreferencesGeneral): void

    //useragent
    setUserAgent(userAgent: string): void

    // proxy
    setProxy(proxy: string): void

    //maxdepth
    setMaxDepth(maxDepth: number): void

    //maxconcurrency
    setMaxConcurrency(maxConcurrency: number): void

    setEnableProcessing(enableProcessing: boolean): void
}

export type GeneralStore = GeneralSlice & GeneralSliceActions;

export const createGeneralSlice: StateCreator<
    GeneralSlice & GeneralSliceActions,
    [["zustand/devtools", never]],
    []
> = (set, get) => {
    return ({
        updateGeneral: (General) => set({General}),
        setUserAgent: (userAgent) => {
            let g = get().General;
            g!.userAgent = userAgent;
            set({General: g}, false, "setUserAgent");
        },
        setProxy: (proxy) => {
            let g = get().General;
            g!.proxyURL = proxy;
            set({General: g}, false, "setProxy");
        },
        setMaxDepth: (maxDepth) => {
            let g = get().General;
            g!.maxDepth = maxDepth;
            set({General: g}, false, "setMaxDepth");
        },
        setMaxConcurrency: (maxConcurrency) => {
            let g = get().General;
            g!.maxConcurrency = maxConcurrency;
            set({General: g}, false, "setMaxConcurrency");
        },
        setEnableProcessing: (enableProcessing) => {
            let g = get().General;
            g!.enableProcessing = enableProcessing;
            set({General: g}, false, "setEnableProcessing");
        },

    });
};