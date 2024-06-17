import {create} from "zustand";
import {createJSONStorage, devtools, persist} from "zustand/middleware";
import {BehaviorStore, createBehaviorSlice} from "./slices/createBehaviorSlice";
import {createGeneralSlice, GeneralStore} from "./slices/createGeneralSlice";


const useAppStore = create<BehaviorStore & GeneralStore>()(devtools(persist((...a) => ({
            ...createBehaviorSlice(...a),
            ...createGeneralSlice(...a),
        }), {
            name: "appStore",
            storage: createJSONStorage(() => localStorage),

        }), {
            name: "appStore", trace: import.meta.env.DEV, enabled: import.meta.env.DEV,
        },
    ),
);


export default useAppStore;