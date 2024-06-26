import {createFileRoute} from '@tanstack/react-router'
import {Box, Flex, IconButton, Skeleton, Text, VStack} from "@chakra-ui/react";
import {CircleArrowLeft, PanelsTopLeft} from "lucide-react";
import {useTranslation} from "react-i18next";
import {useQuery} from "@tanstack/react-query";
import {SearchService} from "../../services/search.ts";
import {ReactNode} from "react";

export const Route = createFileRoute('/queue/statistics')({
    component: ShowStatistics
})

function ShowStatistics() {
    const {t} = useTranslation();

    const {data, isLoading} = useQuery({
        queryFn: async () => SearchService.GetStatistics(),
        queryKey: ['queue', 'statistics', 'get'],
        staleTime: 1000 * 60 * 5, // 5 minutes
        refetchOnWindowFocus: true,
        refetchIntervalInBackground: true,
    })
    return (
        <Box maxH='90vh'>
            <Text fontSize='6xl'>
                <IconButton size={'lg'} aria-label={t('btn.back')} icon={<CircleArrowLeft/>}
                            onClick={() => history.back()}/>
                {t('statistics_list')}</Text>
            <Flex direction={'column'}>
                <Skeleton isLoaded={!isLoading}>
                    <StatCardV2 icon={<PanelsTopLeft size={'100%'}/>} title={'Paginas Indexadas'}
                                value={data?.total_pages}/>
                </Skeleton>
            </Flex>
        </Box>
    )
}

interface StatCardProps {
    icon: ReactNode,
    title: string;
    description?: string;
    value?: string | number;
}

export function StatCardV2({icon, title, description, value}: StatCardProps) {
    return (
        <Box
            display="flex"
            borderWidth="1px"
            borderRadius="lg"
            overflow="hidden"
            width="100%"
            maxW="400px"
            boxShadow="md"
        >
            <Box flex="1">
                {icon}
            </Box>
            <VStack flex="2" p={4} align="start">
                <Text fontWeight="bold" fontSize="xl">
                    {title}
                </Text>
                {description && (
                    <Text fontSize="md" color="gray.500">
                        {description}
                    </Text>
                )}
                <Text fontSize="2xl" color="teal.500">
                    {value}
                </Text>
            </VStack>
        </Box>
    );
};