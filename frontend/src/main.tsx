import React from 'react'
import ReactDOM from 'react-dom/client'

import {ChakraProvider} from "@chakra-ui/react";
import {QueryClient, QueryClientProvider} from '@tanstack/react-query';
import {createMemoryHistory, createRouter, RouterProvider} from '@tanstack/react-router';
import {routeTree} from './routeTree.gen';


// Create a QueryClient instance
const queryClient = new QueryClient();

//ROTAS
// const hashHistory = createHashHistory();
const memoryHistory = createMemoryHistory({
    initialEntries: ['/'], // Pass your initial url
})
const router = createRouter({
    routeTree,
    history: memoryHistory,
    context: {
        queryClient,
    },
    defaultPreload: 'intent',
    // Since we're using React Query, we don't want loader calls to ever be stale
    // This will ensure that the loader is always called when the route is preloaded or visited
    defaultPreloadStaleTime: 0,
})

// Register the router instance for type safety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <ChakraProvider>
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router}/>
            </QueryClientProvider>
        </ChakraProvider>
    </React.StrictMode>,
)
