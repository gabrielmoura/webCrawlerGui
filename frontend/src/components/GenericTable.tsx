import {useTranslation} from "react-i18next";
import {useState} from "react";
import {
    flexRender,
    getCoreRowModel,
    getPaginationRowModel,
    PaginationState,
    useReactTable
} from "@tanstack/react-table";
import {Center, Flex, IconButton, Table, TableContainer, Tbody, Td, Text, Th, Thead, Tr} from "@chakra-ui/react";
import {ArrowLeft, ArrowLeftToLine, ArrowRight, ArrowRightToLine} from "lucide-react";

interface GenericTableProps {
    data: Array<any>;
    columns: Array<any>;
    pageSize?: number;
}

export function GenericTable({data, columns, pageSize = 10}: GenericTableProps) {
    const {t} = useTranslation();

    const [pagination, setPagination] = useState<PaginationState>({
        pageIndex: 0,
        pageSize: pageSize,
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

    return (<TableContainer>
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