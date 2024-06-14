import { createRootRouteWithContext, Outlet, } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { QueryClient } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { FRootLayout } from '../theme/layout'
import {useEffect} from "react";
import {types} from "../../wailsjs/go/models.ts";
import useAppStore from "../store/appStore.ts";
import {GetPreferences} from "../../wailsjs/go/services/ConfigService";

export const Route = createRootRouteWithContext<{
    queryClient: QueryClient
}>()({
    component: RootComponent,
})

function RootComponent() {
    const upB = useAppStore(s => s.updateBehavior)
    const upG = useAppStore(s => s.updateGeneral)
    useEffect(() => {
        // EventsOn('window_changed', (event: any) => {
        //     console.log('window_changed', event)
        // })

        GetPreferences().then((resp: types.JSResp) => {
                console.log('GetPreferences', resp)
                if (resp.success) {
                    let d = resp.data as types.Preferences
                    console.log('GetPreferences', d.behavior)
                    d.behavior.welcomed = true
                    upB(d.behavior)
                    upG(d.general)
                }
            }
        ).catch((err: any) => {
            console.error('GetPreferences', err)
        })
        // Environment().then(env => console.log('Environment', env))
        // Info().then(info => console.log('Info', info))
    }, [])
    return (
        <>
            <FRootLayout>
                <Outlet />
            </FRootLayout>
            {import.meta.env.DEV ?
                <>
                    <ReactQueryDevtools buttonPosition="bottom-left" />
                    <TanStackRouterDevtools position="bottom-right" />
                </>
                : null}
        </>
    )
}