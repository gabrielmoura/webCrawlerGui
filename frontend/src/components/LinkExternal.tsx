import {IconButton, Link, Stack, Text, Tooltip} from "@chakra-ui/react";
import {resumeStringV2} from "../util/helper.ts";
import {ClipboardCopy} from "lucide-react";
import {BrowserOpenURL, ClipboardSetText} from "../../wailsjs/runtime";
import {useTranslation} from "react-i18next";

interface LinkExternalProps {
    url: string;
    maxLength?: number;
}

export function LinkExternal({url, maxLength = 80}: LinkExternalProps) {
    const {t} = useTranslation();
    function handleClipboard() {
        ClipboardSetText(url).catch(console.error);
    }

    return <Stack direction='row'>
        <Tooltip label={t('btn.access')}>
            <Link href={url} isExternal={true} onClick={() => BrowserOpenURL(url)}>
                <Text>{resumeStringV2(url, maxLength)}</Text>
            </Link>
        </Tooltip>
        <Tooltip label={t('btn.copyToClipboard')}>
            <IconButton aria-label={t('btn.copyToClipboard')} onClick={() => handleClipboard()} icon={<ClipboardCopy/>}/>
        </Tooltip>
    </Stack>

}