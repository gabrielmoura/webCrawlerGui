import {IconButton, Tooltip} from "@chakra-ui/react";
import {CirclePause, CirclePlay} from "lucide-react";
import {useMutation} from "@tanstack/react-query";
import useAppStore from "../store/appStore.ts";
import {QueueService} from "../services/queue.ts";
import {useTranslation} from "react-i18next";

export function PlayCrawler() {
    const {t} = useTranslation();
    const enableProcessing = useAppStore(s => s.General?.enableProcessing)
    const setEnableProcessing = useAppStore(s => s.setEnableProcessing)

    const mutate = useMutation({
        mutationKey: ['toggleCrawler'],
        mutationFn: async () => {
            QueueService.toggleCrawling(!!enableProcessing).then(() => {
                setEnableProcessing(!enableProcessing)
            })
        },
    })
    return (
        <Tooltip label={enableProcessing ? t('btn.pause') : t('btn.start')}>
            <IconButton
                aria-label={t('btn.start')}
                icon={enableProcessing ? <CirclePause/> : <CirclePlay/>}
                variant={'outline'}
                colorScheme={'green'}
                onClick={() => mutate.mutate()}
            />
        </Tooltip>
    )
}