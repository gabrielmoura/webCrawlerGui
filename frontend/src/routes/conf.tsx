import {createFileRoute} from '@tanstack/react-router'
import {Box, Flex, Text} from "@chakra-ui/react";
import useAppStore from "../store/appStore";


export const Route = createFileRoute('/conf')({
    component: Configuration
})

function Configuration() {
    const general = useAppStore(state => state.General)

    return (
        <Box>
            <Text fontSize='6xl'>Configuration</Text>
            <Flex direction={'column'}>
                <Text fontSize='4xl'>General</Text>

                {general && (
                    <Flex direction={'column'}>
                        <Text fontSize='2xl'>User Agent</Text>
                        <Text>{general.userAgent}</Text>

                        <Text fontSize='2xl'>Proxy</Text>
                        <Text fontSize='xl'>Proxy Habilitado: {general?.proxyEnabled ? "Sim" : "Não"}</Text>
                        <Text>{general?.proxyURL}</Text>

                        <Text fontSize='2xl'>Max Depth</Text>
                        <Text>{general.maxDepth}</Text>
                        <Text fontSize='2xl'>Max Concurrency</Text>
                        <Text>{general.maxConcurrency}</Text>
                    </Flex>
                )}

            </Flex>
        </Box>
    )
}