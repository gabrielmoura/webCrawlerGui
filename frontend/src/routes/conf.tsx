import {createFileRoute} from '@tanstack/react-router';
import {Button, Flex, IconButton, InputGroup, Stack, Text, Tooltip} from "@chakra-ui/react";
import {useEffect, useState} from "react";
import useAppStore from "../store/appStore";
import {types} from "../../wailsjs/go/models.ts";
import {ConfigService} from "../services/config.ts";
import {useMutation} from "@tanstack/react-query";
import {HardDriveDownload} from "lucide-react";
import {BackupService} from "../services/backup.ts";
import {InputConfig} from "../theme/InputConfig.tsx";
import {SwitchConfig} from "../theme/SwitchConfig.tsx";

export const Route = createFileRoute('/conf')({
    component: Configuration,
});

function Configuration() {
    const [preferences, setPreferences] = useState<types.Preferences | undefined>(undefined);
    const updateGeneral = useAppStore((state) => state.updateGeneral);
    const updateBehavior = useAppStore((state) => state.updateBehavior);

    useEffect(() => {
        ConfigService.Get()
            .then((resp: types.Preferences) => {
                console.log('GetPreferences', resp);
                updateGeneral(resp.general);
                updateBehavior(resp.behavior);
                setPreferences(resp);
            })
            .catch((err: any) => {
                console.error('GetPreferences', err);
            });
    }, []);

    return (
        <Stack direction='column' spacing={10}>
            <Text fontSize='6xl'>Configuração</Text>
            <Flex direction='column'>
                {preferences && <StackConfig pref={preferences}/>}
            </Flex>
            <Flex direction='column'>
                <ExportData/>
            </Flex>
        </Stack>
    );
}

interface StackConfigProps {
    pref: types.Preferences;
}

function StackConfig({pref}: StackConfigProps) {
    const [enabledProxy, setProxy] = useState(pref.general.proxyEnabled);
    const [proxyUrl, setProxyUrl] = useState(pref.general.proxyURL);
    const [userAgent, setUserAgent] = useState(pref.general.userAgent);
    const [maxDepth, setMaxDepth] = useState(pref.general.maxDepth);
    const [maxConcurrency, setMaxConcurrency] = useState(pref.general.maxConcurrency);

    const mutate = useMutation({
        mutationKey: ['updateGeneral', {pref}],
        // @ts-expect-error - types are not correct
        mutationFn: (preferences: types.Preferences) => ConfigService.SaveAllPreferences(preferences),
    })

    function handleSubmit() {
        const data = {
            behavior: pref.behavior,
            general: {
                ...pref.general,
                proxyEnabled: enabledProxy,
                proxyURL: proxyUrl,
                userAgent: userAgent,
                maxDepth: maxDepth,
                maxConcurrency: maxConcurrency
            }
        } as types.Preferences;
        mutate.mutate(data)
    }

    return (
        <Stack spacing={5}>
            <InputGroup>
                <InputConfig label='User Agent' value={userAgent} onChange={setUserAgent}/>
            </InputGroup>

            <InputGroup>
                <Stack direction='column' spacing={4}>
                    <Text fontSize='2xl'>Proxy</Text>
                    <Stack direction='row' spacing={4} align='center'>
                        <SwitchConfig label='Enable Proxy' value={enabledProxy} onChange={setProxy}
                                      name='enable-proxy'/>
                        {enabledProxy && <InputConfig label='Proxy URL' value={proxyUrl} onChange={setProxyUrl}/>}
                    </Stack>
                </Stack>
            </InputGroup>

            <InputGroup>
                <InputConfig label='Max Depth' value={maxDepth} type='number' onChange={setMaxDepth}/>
            </InputGroup>

            <InputGroup>
                <InputConfig label='Max Concurrency' value={maxConcurrency} type='number'
                             onChange={setMaxConcurrency}/>

            </InputGroup>
            <Stack direction='row' spacing={4} align='center'>
                <Button
                    onClick={() => handleSubmit()}
                    loadingText='Loading'
                    colorScheme='teal'
                    variant='outline'
                    spinnerPlacement='start'
                >
                    Atualizar
                </Button>
            </Stack>
        </Stack>
    );
}


function ExportData() {
    const mutExportData = useMutation({
        mutationKey: ['ExportData'],
        mutationFn: () => BackupService.Export(),
    })
    return <Stack direction='row'>
        <Text fontSize='2xl' maxW='50%'>Exportar Dados</Text>
        <Tooltip label='Exportar dados'>
            <IconButton
                icon={<HardDriveDownload/>}
                aria-label='Export Data'
                onClick={() => mutExportData.mutate()}
            />
        </Tooltip>
    </Stack>
}