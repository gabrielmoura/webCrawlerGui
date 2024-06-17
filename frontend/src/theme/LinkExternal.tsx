import {IconButton, Link, Stack, Text, Tooltip} from "@chakra-ui/react";
import {resumeStringV2} from "../util/helper.ts";
import {ClipboardCopy} from "lucide-react";
import {BrowserOpenURL, ClipboardSetText} from "../../wailsjs/runtime";

interface LinkExternalProps {
    url: string;
    maxLength?: number;
}

export function LinkExternal({url, maxLength = 80}: LinkExternalProps) {
    function handleClipboard() {
        ClipboardSetText(url).catch(console.error);
    }

    return <Stack direction='row'>
        <Tooltip label='Acessar'>
            <Link href={url} isExternal={true} onClick={() => BrowserOpenURL(url)}>
                <Text>{resumeStringV2(url, maxLength)}</Text>
            </Link>
        </Tooltip>
        <Tooltip label='Copiar para a área de transferência'>
            <IconButton aria-label='Copy to Clipboard' onClick={() => handleClipboard()} icon={<ClipboardCopy/>}/>
        </Tooltip>
    </Stack>

}