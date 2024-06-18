import {createFileRoute} from "@tanstack/react-router";
import {Button, Flex, InputGroup, Stack, Text, useToast,} from "@chakra-ui/react";
import {useEffect, useState} from "react";
import useAppStore from "../store/appStore";
import {types} from "../../wailsjs/go/models.ts";
import {ConfigService} from "../services/config.ts";
import {useMutation} from "@tanstack/react-query";
import {InputConfig} from "../theme/InputConfig.tsx";
import {SwitchConfig} from "../theme/SwitchConfig.tsx";
import {TagInput} from "../theme/TagInput.tsx";
import {ExportData} from "../theme/ExportData.tsx";


export const Route = createFileRoute("/conf")({
    component: ConfigurationPage,
});

function ConfigurationPage() {
    const [preferences, setPreferences] = useState<types.Preferences | undefined>(
        undefined
    );
    const updateGeneral = useAppStore((state) => state.updateGeneral);
    const updateBehavior = useAppStore((state) => state.updateBehavior);

    useEffect(() => {
        ConfigService.Get()
            .then((resp: types.Preferences) => {
                console.log("GetPreferences", resp);
                updateGeneral(resp.general);
                updateBehavior(resp.behavior);
                setPreferences(resp);
            })
            .catch((err: any) => {
                console.error("GetPreferences", err);
            });
    }, []);

    return (
        <Stack direction="column" spacing={10}>
            <Text fontSize="6xl">Configuração</Text>
            <Flex direction="column">
                {preferences && <StackConfig pref={preferences}/>}
            </Flex>
            <Flex direction="column">
                <ExportData/>
            </Flex>
        </Stack>
    );
}

interface StackConfigProps {
    pref: types.Preferences;
}

function StackConfig({pref}: StackConfigProps) {
    const toast = useToast()

    const [enabledProxy, setProxy] = useState(pref.general.proxyEnabled);
    const [proxyUrl, setProxyUrl] = useState(pref.general.proxyURL);
    const [userAgent, setUserAgent] = useState(pref.general.userAgent);
    const [maxDepth, setMaxDepth] = useState(pref.general.maxDepth);
    const [maxConcurrency, setMaxConcurrency] = useState(
        pref.general.maxConcurrency
    );
    const [ignoreLocal, setIgnoreLocal] = useState<boolean>(pref.general.ignoreLocal);
    const [tlds, setTld] = useState<Array<string>>(pref.general.tlds);
    const regexTld = new RegExp("\\.[a-zA-Z][a-zA-Z0-9]{1,4}\\b");

    const mutate = useMutation({
        mutationKey: ["updateGeneral", {pref}],
        mutationFn: async (preferences: types.Preferences) =>
            ConfigService.SaveAllPreferences(preferences),
        onSuccess: (msg) => {
            toast({
                title: msg,
                status: 'success',
                isClosable: true,
                position: 'bottom-right',
            })
        },
        onError: (msg: string) => {
            toast({
                title: msg,
                status: 'error',
                isClosable: true,
                position: 'bottom-right',
            })
        }
    });

    function handleSubmit() {
        const data = {
            behavior: pref.behavior,
            general: {
                ...pref.general,
                proxyEnabled: enabledProxy,
                ignoreLocal: ignoreLocal,
                proxyURL: proxyUrl,
                userAgent: userAgent,
                maxDepth: maxDepth,
                maxConcurrency: maxConcurrency,
                tlds: tlds,
            },
        } as types.Preferences;
        mutate.mutate(data);
    }

    return (
        <Stack spacing={5}>
            <InputGroup>
                <InputConfig
                    label="User Agent"
                    value={userAgent}
                    onChange={setUserAgent}
                />
            </InputGroup>


            <Stack direction="row" spacing={4} align="center" h={"5rem"}>
                <SwitchConfig
                    label="Enable Proxy"
                    value={enabledProxy}
                    onChange={setProxy}
                    name="enable-proxy"
                />
                <SwitchConfig
                    label="Ignore Local"
                    value={ignoreLocal}
                    onChange={setIgnoreLocal}
                    name="enable-local"
                />
                {enabledProxy && (
                    <InputConfig
                        label="Proxy URL"
                        value={proxyUrl}
                        onChange={setProxyUrl}
                    />
                )}

            </Stack>


            <Stack direction={"row"} spacing={4}>
                <InputGroup>
                    <InputConfig
                        label="Max Depth"
                        value={maxDepth}
                        type="number"
                        onChange={setMaxDepth}
                    />
                </InputGroup>

                <InputGroup>
                    <InputConfig
                        label="Max Concurrency"
                        value={maxConcurrency}
                        type="number"
                        onChange={setMaxConcurrency}
                    />
                </InputGroup>
            </Stack>

            <Stack spacing={4}>
                <Text fontSize="2xl">TLDs</Text>
                <TagInput
                    setTags={setTld}
                    tags={tlds}
                    regex={regexTld}
                    placeholder="Inclua um tld EX: .com"
                />
            </Stack>

            <Stack direction="row" spacing={4} align="center">
                <Button
                    onClick={() => handleSubmit()}
                    isLoading={mutate.isPending}
                    loadingText="Loading"
                    colorScheme="teal"
                    variant="outline"
                    spinnerPlacement="start"
                >
                    Atualizar
                </Button>
            </Stack>
        </Stack>
    );
}
