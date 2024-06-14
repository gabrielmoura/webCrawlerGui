import {createFileRoute} from '@tanstack/react-router'
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
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {useState} from 'react';
import {ArrowLeft, ArrowLeftToLine, ArrowRight, ArrowRightToLine, CircleX, Plus} from "lucide-react";
import {QueueService, URL} from "../services/queue";
import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    getPaginationRowModel,
    PaginationState,
    useReactTable
} from '@tanstack/react-table';

export const Route = createFileRoute('/queue')({
    component: ShowQueueList
})

function ShowQueueList() {
    const [url, setUrl] = useState<string>('')
    const client = useQueryClient()
    const {data} = useQuery({
        queryFn: async () => QueueService.getAllQueue(),
        queryKey: ['queue', 'get'],
        initialData: []
    })

    const mutateCreate = useMutation({
        mutationKey: ['queue', 'set'],
        mutationFn: async (url: string) => QueueService.addToQueue(url),
        onSuccess: async () => {
            await client.invalidateQueries({queryKey: ['queue', 'get']})
        }
    })

    const mutateDelete = useMutation({
        mutationKey: ['queue', 'delete'],
        mutationFn: async (url: string) => QueueService.deleteQueue(url),
        onSuccess: async () => {
            await client.invalidateQueries({queryKey: ['queue', 'get']})
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

    const columnHelper = createColumnHelper<URL>();
    const columns = [
        columnHelper.accessor('url', {
            cell: info => <span className='font-bold text-blue-500 hover:underline'>{info.getValue()}</span>,
            footer: info => info.column.id,
        }),
        columnHelper.accessor('depth', {
            id: 'group',
            cell: info => <i>{info.getValue()}</i>,
            header: () => <span>Depth</span>,
            footer: info => info.column.id,
        }),
        // Actions
        columnHelper.accessor('url', {
            header: 'Actions',
            footer: 'Actions',
            cell: info => (
                <IconButton
                    onClick={() => handleDeleteFromQueue(info.getValue())}
                    aria-label='Delete'
                    icon={<CircleX/>}
                />
            ),
        }),
    ];
    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: 10,
    })

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        state: {
            pagination,
        },
        onPaginationChange: setPagination,
    });

    return (
        <Box maxH='90vh'>
            <Text fontSize='6xl'>Queue List</Text>
            <Flex direction={'column'}>
                <Center>
                    <InputGroup>
                        <InputLeftAddon>Url</InputLeftAddon>
                        <Input type='text' placeholder='URL to scrawler' onChange={({target}) => setUrl(target.value)}/>
                    </InputGroup>
                    <Button onClick={() => handleAddToQueue()}>
                        <Tooltip label='Incluir'>
                            <Plus/>
                        </Tooltip>
                    </Button>
                </Center>
                <TableContainer>
                    <Table variant='striped'>
                        <Thead>
                            {table.getHeaderGroups().map(headerGroup => (
                                <Tr key={headerGroup.id}>
                                    {headerGroup.headers.map(header => (
                                        <Th key={header.id}>
                                            {flexRender(
                                                header.column.columnDef.header,
                                                header.getContext()
                                            )}
                                        </Th>
                                    ))}
                                </Tr>
                            ))}
                        </Thead>
                        <Tbody>
                            {table.getRowModel().rows.map(row => (
                                <Tr key={row.id}>
                                    {row.getVisibleCells().map(cell => (
                                        <Td key={cell.id}>
                                            {flexRender(
                                                cell.column.columnDef.cell,
                                                cell.getContext()
                                            )}
                                        </Td>
                                    ))}
                                </Tr>
                            ))}
                        </Tbody>
                    </Table>
                    <Flex gap={5}>
                        <IconButton
                            className="border rounded p-1"
                            onClick={() => table.firstPage()}
                            disabled={!table.getCanPreviousPage()}
                            aria-label='Primeiro'
                            icon={<ArrowLeftToLine/>}
                        />
                        <IconButton
                            className="border rounded p-1"
                            onClick={() => table.previousPage()}
                            disabled={!table.getCanPreviousPage()}
                            aria-label='Voltar'
                            icon={<ArrowLeft/>}
                        />
                        <Text w={170} h={10}>
                            {table.getState().pagination.pageIndex + 1} of{' '}
                            {table.getPageCount().toLocaleString()} of {table.getRowCount().toLocaleString()} Rows
                        </Text>
                        <IconButton
                            className="border rounded p-1"
                            onClick={() => table.nextPage()}
                            disabled={!table.getCanNextPage()}
                            icon={<ArrowRight/>}
                            aria-label='Avançar'
                        />
                        <IconButton
                            className="border rounded p-1"
                            onClick={() => table.lastPage()}
                            disabled={!table.getCanNextPage()}
                            icon={<ArrowRightToLine/>}
                            aria-label='Último'
                        />
                    </Flex>

                </TableContainer>
            </Flex>
        </Box>
    );
}
