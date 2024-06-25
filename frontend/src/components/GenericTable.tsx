import {useTranslation} from "react-i18next";
import {CSSProperties, useState} from "react";
import {
    flexRender,
    getCoreRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    PaginationState,
    SortingState,
    useReactTable
} from "@tanstack/react-table";
import {Center, Flex, IconButton, Table, TableContainer, Tbody, Td, Text, Th, Thead, Tr} from "@chakra-ui/react";
import {ArrowDownZA, ArrowLeft, ArrowLeftToLine, ArrowRight, ArrowRightToLine, ArrowUpAZ} from "lucide-react";

interface GenericTableProps {
    data: Array<any>;
    columns: Array<any>;
    pageSize?: number;
}

const TableCss: Record<string, CSSProperties> = {
    cursor_pointer: {
        cursor: 'pointer',
    },
    cursor_none: {},
}

export function GenericTable({data, columns, pageSize = 10}: GenericTableProps) {
    const {t} = useTranslation();

    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: pageSize,
    })
    const [sorting, setSorting] = useState<SortingState>([])

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        getSortedRowModel: getSortedRowModel(),
        state: {
            pagination,
            sorting
        },
        onPaginationChange: setPagination,
        onSortingChange: setSorting,
    });

    return (<TableContainer>
        <Table variant='striped'>
            <Thead>
                {table.getHeaderGroups().map(headerGroup => (
                    <Tr key={headerGroup.id}>
                        {headerGroup.headers.map(header => (
                            <Th key={header.id}
                                style={header.column.getCanSort() ? TableCss.cursor_pointer : TableCss.cursor_none}
                                {...{
                                    onClick: header.column.getToggleSortingHandler(),
                                }}
                            >
                                <Flex direction={'row'}>
                                    {flexRender(
                                        header.column.columnDef.header,
                                        header.getContext()
                                    )}

                                    {{
                                        asc: <><ArrowUpAZ size={20}/></>,
                                        desc: <><ArrowDownZA size={20}/></>,
                                    }[header.column.getIsSorted() as string] ?? null}
                                </Flex>
                            </Th>
                        ))}
                    </Tr>
                ))}
            </Thead>
            <Tbody>
                {table.getRowModel().rows.map(row => (
                    <Tr key={row.id}>
                        {row.getVisibleCells().map(cell => (
                            <Td key={cell.id} maxW={'70vw'}
                                style={{
                                    wordBreak: 'break-word',
                                    whiteSpace: 'normal',
                                }}
                            >
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
                    isDisabled={!table.getCanPreviousPage()}
                    aria-label={t('btn.first')}
                    icon={<ArrowLeftToLine/>}
                />
                <IconButton
                    className="border rounded p-1"
                    onClick={() => table.previousPage()}
                    isDisabled={!table.getCanPreviousPage()}
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
                    isDisabled={!table.getCanNextPage()}
                    icon={<ArrowRight/>}
                    aria-label={t('btn.next')}
                />
                <IconButton
                    className="border rounded p-1"
                    onClick={() => table.lastPage()}
                    isDisabled={!table.getCanNextPage()}
                    icon={<ArrowRightToLine/>}
                    aria-label={t('btn.last')}
                />
            </Flex>
        </Center>

    </TableContainer>)
}