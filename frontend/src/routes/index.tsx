import {useMutation} from '@tanstack/react-query'
import {createFileRoute} from '@tanstack/react-router'
import {useState} from 'react'
import {SearchService} from '../services/search'
import {Box, Center, Flex, IconButton, Input, InputGroup, InputLeftAddon, Tooltip} from '@chakra-ui/react'
import {Search} from 'lucide-react'


export const Route = createFileRoute('/')({
    component: SearchPage,
})


function SearchPage() {
    const [searchString, setSearch] = useState<string>()
    const [data, setData] = useState<any>([])
    const mutSearch = useMutation({
        mutationKey: ['search', {searchString}],
        mutationFn: (query: string) => SearchService.search(query),
        onSuccess: (data) => {
            setData(data)
        }
    })

    function HandleSearch() {
        if (searchString && searchString.length >= 3) {
            mutSearch.mutate(searchString)
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
                        />
                    </InputGroup>
                    <Tooltip label='Buscar'>
                        <IconButton aria-label='Search' icon={<Search/>} onClick={() => HandleSearch()}/>
                    </Tooltip>
                </Center>
                {data?.map((item: any) => (
                    <tbody>
                    <tr>
                        <td>{item?.url}</td>
                        <td>{item?.title}</td>
                        <td>{item?.description}</td>
                    </tr>
                    </tbody>
                ))
                }

            </Flex>
        </Box>
    )
}
