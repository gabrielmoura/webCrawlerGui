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
    Text,
    Tooltip,
    useToast
} from "@chakra-ui/react";
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {useState} from 'react';
import {CircleX, Plus} from "lucide-react";
import {QueueService, URL} from "../services/queue";
import {createColumnHelper} from '@tanstack/react-table';
import {useTranslation} from "react-i18next";
import useAppStore from "../store/appStore.ts";
import {AddHostsTxt} from "../components/AddHostsTxt.tsx";
import {onEnter} from "../util/helper.ts";
import {GenericTable} from "../components/GenericTable.tsx";

export const Route = createFileRoute('/queue')({
    component: ShowQueueList
})

function ShowQueueList() {
    const importsEnabled = useAppStore<boolean>(s => s.General!.enableImportHosts)
    const {t} = useTranslation();
    const toast = useToast()
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
        onSuccess: (msg) => {
            setUrl('')
            toast({
                title: t(`msg.${msg}`),
                status: 'success',
                duration: 9000,
                isClosable: true,
            })
            client.refetchQueries({queryKey: ['queue', 'get']}).catch(console.error)

        },
        onError: (msg) => {
            toast({
                title: t(`msg.${msg}`),
                status: 'error',
                duration: 9000,
                isClosable: true,
            })
        }
    })

    const mutateDelete = useMutation({
        mutationKey: ['queue', 'delete'],
        mutationFn: async (url: string) => QueueService.deleteQueue(url),
        onSuccess: async (msg) => {
            await client.invalidateQueries({queryKey: ['queue', 'get']})
            toast({
                title: t(`msg.${msg}`),
                status: 'success',
                duration: 9000,
                isClosable: true,
            })
        }, onError: (msg) => {
            toast({
                title: t(`msg.${msg}`),
                status: 'error',
                duration: 9000,
                isClosable: true,
            })
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


    return (
        <Box maxH='90vh'>
            <Text fontSize='6xl'>{t('queue_list')}</Text>
            <Flex direction={'column'}>
                <Center my={'0.5rem'} gap={1}>
                    <InputGroup>
                        <InputLeftAddon>Url</InputLeftAddon>
                        <Input type='text' placeholder={t('placeholder.url')}
                               value={url}
                               onChange={({target}) => setUrl(target.value)}
                               onKeyDown={e => onEnter(e, handleAddToQueue)}
                        />
                    </InputGroup>
                    <Button onClick={() => handleAddToQueue()} isLoading={mutateCreate.isPending}>
                        <Tooltip label={t('btn.include')}>
                            <Plus/>
                        </Tooltip>
                    </Button>
                </Center>

                {importsEnabled ? <AddHostsTxt/> : null}
                {data && data.length > 0 ? <GenericTable data={data} columns={columns}/> : null}
            </Flex>
        </Box>
    );
}