import {IconButton, Stack, Text, Tooltip, useToast} from "@chakra-ui/react";
import {useTranslation} from "react-i18next";
import {useMutation} from "@tanstack/react-query";
import {BackupService} from "../services/backup.ts";
import {BookPlus} from "lucide-react";

export function ImportQueue(){
    const toast = useToast()
    const {t} = useTranslation();

    const mutImportQueue = useMutation({
        mutationKey: ["ImportQueue"],
        mutationFn: () => BackupService.ImportQueue(),
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
                {t('import_label')}
            </Text>
            <Tooltip label={t('import_label')}>
                <IconButton
                    icon={<BookPlus/>}
                    aria-label={t('import_label')}
                    onClick={() => mutImportQueue.mutate()}
                    isLoading={mutImportQueue.isPending}
                />
            </Tooltip>
        </Stack>
    );
}