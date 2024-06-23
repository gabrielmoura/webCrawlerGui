import {createRootRouteWithContext, Outlet,} from '@tanstack/react-router'
import {TanStackRouterDevtools} from '@tanstack/router-devtools'
import {QueryClient} from '@tanstack/react-query'
import {ReactQueryDevtools} from '@tanstack/react-query-devtools'
import {useEffect} from "react";
import {types} from "../../wailsjs/go/models.ts";
import useAppStore from "../store/appStore.ts";
import {ConfigService} from "../services/config.ts";
import {FRootLayout} from "../components/Layout.tsx";

export const Route = createRootRouteWithContext<{
    queryClient: QueryClient
}>()({
    component: RootComponent,
})

function RootComponent() {
    const upB = useAppStore(s => s.updateBehavior)
    const upG = useAppStore(s => s.updateGeneral)
    useEffect(() => {
        ConfigService.Get().then((resp: types.Preferences) => {
                console.log('GetPreferences', resp)
                upB(resp.behavior)
                upG(resp.general)
            }
        ).catch((err: any) => {
            console.error('GetPreferences', err)
        })
    }, [])
    return (
        <>
            <FRootLayout>
                <Outlet/>
            </FRootLayout>
            {import.meta.env.DEV ?
                <>
                    <ReactQueryDevtools buttonPosition="bottom-left"/>
                    <TanStackRouterDevtools position="bottom-right"/>
                </>
                : null}
        </>
    )
}