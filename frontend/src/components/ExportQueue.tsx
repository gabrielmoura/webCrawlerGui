import {IconButton, Stack, Text, Tooltip, useToast} from "@chakra-ui/react";
import {useTranslation} from "react-i18next";
import {useMutation} from "@tanstack/react-query";
import {BackupService} from "../services/backup.ts";
import {BookDown} from "lucide-react";

export function ExportQueue() {
    const toast = useToast()
    const {t} = useTranslation();

    const mutExportQueue = useMutation({
        mutationKey: ["ExportQueue"],
        mutationFn: () => BackupService.ExportQueue(),
        onSuccess: (msg) => {
            toast({
                title: t(`msg.${msg}`),
                status: 'success',
                isClosable: true,
                position: 'bottom-right',
            })
        },
        onError: (msg: string) => {
            toast({
                title: t(`msg.${msg}`),
                status: 'error',
                isClosable: true,
                position: 'bottom-right',
            })
        }
    });


    return (
        <Stack direction="row">
            <Text fontSize="2xl" >
                {t('export_label')}
            </Text>
            <Tooltip label={t('export_label')}>
                <IconButton
                    icon={<BookDown/>}
                    aria-label={t('export_label')}
                    onClick={() => mutExportQueue.mutate()}
                    isLoading={mutExportQueue.isPending}
                />
            </Tooltip>
        </Stack>
    );
}