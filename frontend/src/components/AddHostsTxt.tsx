import {Button, Center, Input, InputGroup, InputLeftAddon, Tooltip, useToast} from "@chakra-ui/react";
import {Plus} from "lucide-react";
import {useState} from "react";
import {QueueService} from "../services/queue.ts";
import {useTranslation} from "react-i18next";

export function AddHostsTxt() {
    const [hosts, setHosts] = useState<string>();
    const toast = useToast()
    const {t} = useTranslation();

    function handleAddHosts() {
        // Add hosts to the hosts file
        if (hosts) {
            QueueService.addHotsTxt(hosts).then((res) => {
                toast({
                    title: "Hosts added",
                    description: res,
                    status: "success",
                    duration: 9000,
                    isClosable: true,
                })
            }).catch((err) => {
                toast({
                    title: "Hosts not added",
                    description: err,
                    status: "error",
                    duration: 9000,
                    isClosable: true,
                })
            })
        }
    }

    return (
        <Center my={'0.5rem'}>
            <InputGroup>
                <InputLeftAddon>Url do Hosts.txt</InputLeftAddon>
                <Input
                    defaultValue={hosts}
                    onChange={({target}) => setHosts(target.value)}
                />
            </InputGroup>
            <Button onClick={() => handleAddHosts()}>
                <Tooltip label={t('btn.include')}>
                    <Plus/>
                </Tooltip>
            </Button>
        </Center>
    )
}