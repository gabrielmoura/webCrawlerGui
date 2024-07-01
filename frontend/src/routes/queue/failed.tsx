import {createFileRoute} from '@tanstack/react-router'
import {Box, Button, Center, Flex, IconButton, Skeleton, Text, Tooltip, useToast} from "@chakra-ui/react";
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {FailedType, QueueService} from "../../services/queue.ts";
import {useTranslation} from "react-i18next";
import {GenericTable} from "../../components/GenericTable.tsx";
import {createColumnHelper} from "@tanstack/react-table";
import {CircleArrowLeft, Trash2, Undo2} from "lucide-react";

export const Route = createFileRoute('/queue/failed')({
    component: QueueFailed,
})

function QueueFailed() {
    const {t} = useTranslation();
    const toast = useToast()
    const client = useQueryClient()
    const {data, isLoading} = useQuery({
        queryFn: async () => QueueService.getAllFailed(),
        queryKey: ['queue', 'failed', 'get'],
        initialData: []
    })
    const mutRemoveAll = useMutation({
        mutationKey: ['queue', 'failed', 'deleteAll'],
        mutationFn: QueueService.deleteAllFailed,
        onSuccess: () => {
            toast({
                title: t('msg.delete_all_success'),
                status: 'success',
                duration: 9000,
                isClosable: true,
            })
            Promise.all([
            client.refetchQueries({queryKey: ['queue', 'get']}),
            client.resetQueries({queryKey: ['queue', 'failed', 'get']})
            ]).catch(console.error)
        }
    })

    const mutAdd = useMutation({
        mutationKey: ['queue', 'add'],
        mutationFn: (url: string) => QueueService.addToQueue(url),
        onSuccess: () => {
            toast({
                title: t('msg.retry_success'),
                status: 'success',
                duration: 9000,
                isClosable: true,
            })
            client.refetchQueries({queryKey: ['queue', 'get']}).catch(console.error)
        }
    })

    const mutRemove = useMutation({
        mutationKey: ['queue', 'failed', 'remove'],
        mutationFn: (data: string) => QueueService.deleteFailed(data),
        onMutate: (url: string) => {
            mutAdd.mutate(url)
            return url
        }
    })

    function handleRemoveAll() {
        mutRemoveAll.mutate()
    }

    function handleRetry(url: string) {
        mutRemove.mutate(url)
    }

    const columnHelper = createColumnHelper<FailedType>();
    const columns = [
        columnHelper.accessor('url', {
            cell: info => <span className='font-bold text-blue-500 hover:underline'>{info.getValue()}</span>,
            footer: info => info.column.id,
            enableSorting: false,
        }),
        columnHelper.accessor('reason', {
            id: 'reason',
            cell: info => <i>{info.getValue()}</i>,
            header: () => <Text>{t('reason')}</Text>,
            footer: info => info.column.id,
            enableSorting: true,
        }),
        // Actions
        columnHelper.accessor('url', {
            header: t('actions'),
            footer: t('actions'),
            cell: info => (
                <Tooltip label={t('btn.retry')}>
                    <IconButton
                        isLoading={mutRemove.isPending}
                        onClick={() => handleRetry(info.getValue())}
                        aria-label={t('btn.retry')}
                        icon={<Undo2/>}
                    />
                </Tooltip>
            ),
            enableSorting: false,
        }),
    ];
    return (
        <Box maxH='90vh'>
            <Text fontSize='6xl'>
                <IconButton size={'lg'} aria-label={t('btn.back')} icon={<CircleArrowLeft/>}
                            onClick={() => history.back()}/>
                {t('failed_list')}</Text>
            <Flex direction={'column'}>
                <Button onClick={() => handleRemoveAll()} isLoading={mutRemoveAll.isPending}>
                    <Center gap={1}>
                        <Text>
                            Remover Todos
                        </Text>
                        <Trash2/>
                    </Center>
                </Button>
                <Skeleton isLoaded={!isLoading}>
                    <GenericTable data={data} columns={columns}/>
                </Skeleton>
            </Flex>
        </Box>
    )
}