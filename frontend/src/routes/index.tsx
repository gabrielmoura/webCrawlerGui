import {useMutation} from '@tanstack/react-query'
import {createFileRoute} from '@tanstack/react-router'
import {useState} from 'react'
import {Page, SearchService} from '../services/search'
import {
    Box,
    Center,
    Flex,
    IconButton,
    Input,
    InputGroup,
    InputLeftAddon,
    Spinner,
    Text,
    Tooltip
} from '@chakra-ui/react'
import {Search} from 'lucide-react'
import {useTranslation} from "react-i18next";
import {LinkExternal} from "../components/LinkExternal.tsx";
import {onEnter} from "../util/helper.ts";
import {createColumnHelper} from "@tanstack/react-table";
import {GenericTable} from "../components/GenericTable.tsx";


export const Route = createFileRoute('/')({
    component: SearchPage,
})


function SearchPage() {
    const {t} = useTranslation()
    const [searchString, setSearch] = useState<string>()
    const [data, setData] = useState<Array<Page>>([])
    const mutSearch = useMutation({
        mutationKey: ['search', {searchString}],
        mutationFn: (query: string) => SearchService.searchWords(query),
        onSuccess: (data) => {
            setData(data)
        }
    })

    function HandleSearch() {
        if (searchString && searchString.length >= 3) {
            mutSearch.mutate(searchString)
        }
    }

    const columnHelper = createColumnHelper<Page>();
    const columns = [
        columnHelper.accessor('url', {
            cell: info => <LinkExternal url={info.getValue()}/>,
            footer: info => info.column.id,
        }),
        columnHelper.accessor('title', {
            id: 'title',
            cell: info => <i>{info.getValue()}</i>,
            header: () => <span>{t('title')}</span>,
            footer: info => info.column.id,
        }),
        columnHelper.accessor('description', {
            id: 'description',
            cell: info => <i>{info.getValue()}</i>,
            header: () => <span>{t('description')}</span>,
            footer: info => info.column.id,
        }),
    ];


    return (
        <Box>
            <Flex gap='2' direction={'column'}>
                <Center gap={1}>
                    <InputGroup>
                        <InputLeftAddon>{t('btn.search')}</InputLeftAddon>
                        <Input type='text' placeholder={t('placeholder.search')}
                               onChange={(e) => setSearch(e.target.value)}
                               onKeyDown={e => onEnter(e, HandleSearch)}
                        />
                    </InputGroup>
                    <Tooltip label={t('btn.search')}>
                        <IconButton aria-label={t('btn.search')} icon={mutSearch.isPending ? <Spinner/> : <Search/>}
                                    onClick={() => HandleSearch()}/>
                    </Tooltip>
                </Center>
                {mutSearch.isSuccess && data?.length > 0 ? <Text>{data?.length} {t('results_found')}</Text> : null}
                {mutSearch.isError && <div>Error</div>}
                {mutSearch.isSuccess && (<GenericTable data={data} columns={columns}/>)}
            </Flex>
        </Box>
    )
}