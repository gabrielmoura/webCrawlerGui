import {IconButton, Stack, Text, Tooltip, useToast} from "@chakra-ui/react";
import {useMutation} from "@tanstack/react-query";
import {BackupService} from "../services/backup.ts";
import {HardDriveDownload} from "lucide-react";

export function ExportData() {
    const toast = useToast()
    const mutExportData = useMutation({
        mutationKey: ["ExportData"],
        mutationFn: () => BackupService.Export(),
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


    return (
        <Stack direction="row">
            <Text fontSize="2xl" maxW="50%">
                Exportar Dados
            </Text>
            <Tooltip label="Exportar dados">
                <IconButton
                    icon={<HardDriveDownload/>}
                    aria-label="Export Data"
                    onClick={() => mutExportData.mutate()}
                />
            </Tooltip>
        </Stack>
    );
}