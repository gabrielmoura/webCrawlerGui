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
import {useTranslation} from "react-i18next";
import {AddHostsTxt} from "../theme/AddHostsTxt.tsx";
import useAppStore from "../store/appStore.ts";

export const Route = createFileRoute('/queue')({
    component: ShowQueueList
})

function ShowQueueList() {
    const importsEnabled = useAppStore<boolean>(s => s.General!.enableImportHosts)
    const {t} = useTranslation();
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
            id: 'depth',
            cell: info => <i>{info.getValue()}</i>,
            header: () => <span>{t('depth')}</span>,
            footer: info => info.column.id,
        }),
        // Actions
        columnHelper.accessor('url', {
            header: t('actions'),
            footer: t('actions'),
            cell: info => (
                <IconButton
                    onClick={() => handleDeleteFromQueue(info.getValue())}
                    aria-label={t('btn.delete')}
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
            <Text fontSize='6xl'>{t('queue_list')}</Text>
            <Flex direction={'column'}>
                <Center my={'0.5rem'}>
                    <InputGroup>
                        <InputLeftAddon>Url</InputLeftAddon>
                        <Input type='text' placeholder={t('placeholder.url')}
                               onChange={({target}) => setUrl(target.value)}/>
                    </InputGroup>
                    <Button onClick={() => handleAddToQueue()}>
                        <Tooltip label={t('btn.include')}>
                            <Plus/>
                        </Tooltip>
                    </Button>
                </Center>

                {importsEnabled ? <AddHostsTxt/> : null}

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
                                        <Td key={cell.id} maxW={'70vw'}>
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
                    <Center>
                        <Flex gap={5}>
                            <IconButton
                                className="border rounded p-1"
                                onClick={() => table.firstPage()}
                                disabled={!table.getCanPreviousPage()}
                                aria-label={t('btn.first')}
                                icon={<ArrowLeftToLine/>}
                            />
                            <IconButton
                                className="border rounded p-1"
                                onClick={() => table.previousPage()}
                                disabled={!table.getCanPreviousPage()}
                                aria-label={t('btn.back')}
                                icon={<ArrowLeft/>}
                            />
                            <Text w={170} h={10}>
                                {table.getState().pagination.pageIndex + 1} {t('of')}{' '}
                                {table.getPageCount().toLocaleString()} {t('of')} {table.getRowCount().toLocaleString()} {t('rows')}
                            </Text>
                            <IconButton
                                className="border rounded p-1"
                                onClick={() => table.nextPage()}
                                disabled={!table.getCanNextPage()}
                                icon={<ArrowRight/>}
                                aria-label={t('btn.next')}
                            />
                            <IconButton
                                className="border rounded p-1"
                                onClick={() => table.lastPage()}
                                disabled={!table.getCanNextPage()}
                                icon={<ArrowRightToLine/>}
                                aria-label={t('btn.last')}
                            />
                        </Flex>
                    </Center>

                </TableContainer>
            </Flex>
        </Box>
    );
}
