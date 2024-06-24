import {IconButton, Stack, Text, Tooltip, useToast} from "@chakra-ui/react";
import {useTranslation} from "react-i18next";
import {useMutation} from "@tanstack/react-query";
import {BackupService} from "../services/backup.ts";
import {HardDriveUpload} from "lucide-react";

export function ImportData() {
    const toast = useToast()
    const {t} = useTranslation();

    const mutImportData = useMutation({
        mutationKey: ["ExportData"],
        mutationFn: () => BackupService.Import(),
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
            <Text fontSize="2xl" maxW="50%">
                {t('import_label')}
            </Text>
            <Tooltip label={t('import_label')}>
                <IconButton
                    icon={<HardDriveUpload/>}
                    aria-label={t('import_label')}
                    onClick={() => mutImportData.mutate()}
                    isLoading={mutImportData.isPending}
                />
            </Tooltip>
        </Stack>
    );
}