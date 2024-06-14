import {IconButton} from "@chakra-ui/react";
import {CirclePause, CirclePlay} from "lucide-react";
import {useMutation} from "@tanstack/react-query";
import useAppStore from "../store/appStore.ts";
import {QueueService} from "../services/queue.ts";

export function PlayCrawler() {
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
        <IconButton
            aria-label={'Play Crawler'}
            icon={enableProcessing ? <CirclePause/> : <CirclePlay/>}
            variant={'outline'}
            colorScheme={'green'}
            onClick={() => mutate.mutate()}
        />
    )
}