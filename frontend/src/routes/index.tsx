import { useMutation } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'
import { SearchService } from '../services/search'
import { Center, Flex, IconButton, Input, InputGroup, InputLeftAddon, Table, Box, Tooltip } from '@chakra-ui/react'
import { Search } from 'lucide-react'


export const Route = createFileRoute('/')({
  component: SearchPage,
})


function SearchPage() {
  const [searchString, setSearch] = useState<string>()
  const mutSearch = useMutation({
    mutationKey: ['search'],
    mutationFn: (query: string) => SearchService.search(query)
  })

  function HandleSearch() {
    if (searchString && searchString.length > 3) {
      mutSearch.mutate(searchString)
    }
  }

  return (
    <Box>
      <Flex gap='2' direction={'column'}>
        <Center >
          <InputGroup>
            <InputLeftAddon>Search</InputLeftAddon>
            <Input type='text' placeholder='Search on scrawler'
              onChange={(e) => setSearch(e.target.value)}
            />
          </InputGroup>
          <Tooltip label='Buscar'>
            <IconButton aria-label='Search' icon={<Search />} onClick={() => HandleSearch()} />
          </Tooltip>
        </Center>
        <Table variant='striped'>
          ola
        </Table>

      </Flex>
    </Box>
  )
}
