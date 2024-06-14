import { createFileRoute } from '@tanstack/react-router'
import {
    Box,
    Button,
    Center,
    Flex,
    IconButton,
    Input,
    InputGroup,
    InputLeftAddon,
    Table,
    TableContainer,
    Tbody,
    Td,
    Text,
    Th,
    Thead,
    Tooltip,
    Tr
} from "@chakra-ui/react";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useState } from 'react';
import { CircleX, Plus } from "lucide-react";
import { QueueService, Url } from "../services/queue";

export const Route = createFileRoute('/queue')({
    component: ShowQueueList
})

function ShowQueueList() {
    const [url, setUrl] = useState<string>()
    const client = useQueryClient()
    const { data } = useQuery({
        queryFn: async () => QueueService.getAllQueue(),
        queryKey: ['queue', 'get'],
        initialData: []
    })
    const mutateCreate = useMutation({
        mutationKey: ['queue', 'set'],
        mutationFn: async (url: string) => QueueService.addToQueue(url),
        onSuccess: async () => {
            await client.invalidateQueries({ queryKey: ['queue', 'get'] })
        }
    })
    const mutateDelete = useMutation({
        mutationKey: ['queue', 'delete'],
        mutationFn: async (url: string) => QueueService.deleteQueue(url),
        onSuccess: async () => {
            await client.invalidateQueries({ queryKey: ['queue', 'get'] })
        }
    })

    function handleAddToQueue() {
        if (url) {
            mutateCreate.mutate(url)
        }
    }

    function handleDeleteFromQueue(url: string) {
        if (url) {
            mutateDelete.mutate(url)
        }
    }

    return (
        <Box>
            <Text fontSize='6xl'>Queue List</Text>
            <Flex direction={'column'}>
                <Center>
                    <InputGroup>
                        <InputLeftAddon>Url</InputLeftAddon>
                        <Input type='text' placeholder='URL to scrawler'
                            onChange={({ target }) => setUrl(target.value)} />
                    </InputGroup>
                    <Button onClick={() => handleAddToQueue()}>
                        <Tooltip label='Incluir'>
                            <Plus />
                        </Tooltip>
                    </Button>
                </Center>
                <TableContainer>
                    <Table variant='striped'>
                        <Thead>
                            <Tr>
                                <Th>URL</Th>
                                <Th>depth</Th>
                            </Tr>
                        </Thead>
                        <Tbody>
                            {data?.map((item: Url) => {
                                return (
                                    <Tr key={item.url}>
                                        <Td>{item.url}</Td>
                                        <Td>{item.depth}</Td>
                                        <Td>
                                            <Tooltip label='Remover'>
                                                <IconButton aria-label='delete' icon={<CircleX />}
                                                    onClick={() => handleDeleteFromQueue(item.url)}
                                                />
                                            </Tooltip>
                                        </Td>
                                    </Tr>
                                )
                            })}
                        </Tbody>
                    </Table>
                </TableContainer>
            </Flex>
        </Box>
    );
}