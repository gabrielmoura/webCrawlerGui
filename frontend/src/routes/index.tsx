import {useMutation} from '@tanstack/react-query'
import {createFileRoute} from '@tanstack/react-router'
import {useState} from 'react'
import {SearchService} from '../services/search'
import {
    Box,
    Center,
    Flex,
    IconButton,
    Input,
    InputGroup,
    InputLeftAddon,
    Spinner,
    Table,
    Tbody,
    Td,
    Thead,
    Tooltip,
    Tr
} from '@chakra-ui/react'
import {Search} from 'lucide-react'
import {LinkExternal} from "../theme/LinkExternal.tsx";


export const Route = createFileRoute('/')({
    component: SearchPage,
})


function SearchPage() {
    const [searchString, setSearch] = useState<string>()
    const [data, setData] = useState<any>([])
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

    function handleEnter(e: any) {
        if (e.key === 'Enter') {
            HandleSearch()
        }
    }


    return (
        <Box>
            <Flex gap='2' direction={'column'}>
                <Center>
                    <InputGroup>
                        <InputLeftAddon>Search</InputLeftAddon>
                        <Input type='text' placeholder='Search on scrawler'
                               onChange={(e) => setSearch(e.target.value)}
                               onKeyDown={handleEnter}
                        />
                    </InputGroup>
                    <Tooltip label='Buscar'>
                        <IconButton aria-label='Search' icon={mutSearch.isPending ? <Spinner/> : <Search/>}
                                    onClick={() => HandleSearch()}/>
                    </Tooltip>
                </Center>
                {mutSearch.isError && <div>Error</div>}
                {mutSearch.isSuccess && (
                    <Table>
                        <Thead>
                            <Tr>
                                <Td>URL</Td>
                                <Td>Título</Td>
                                <Td>Descrição</Td>
                            </Tr>
                        </Thead>
                        {data?.map((item: any) => (
                            <Tbody>
                                <Tr>
                                    <Td>
                                        <LinkExternal url={item?.url}/>
                                    </Td>
                                    <Td>{item?.title}</Td>
                                    <Td>{item?.description}</Td>
                                </Tr>
                            </Tbody>
                        ))
                        }
                    </Table>
                )}


            </Flex>
        </Box>
    )
}
