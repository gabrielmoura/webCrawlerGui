import {ButtonGroup, Flex, IconButton, Spacer, Text, Tooltip} from "@chakra-ui/react";
import {EyeOff, Menu} from "lucide-react";
import {PlayCrawler} from "./PlayCrawler.tsx";
import {ToggleTheme} from "./ToggleTheme.tsx";
import {ConfigService} from "../services/config.ts";
import {useTranslation} from "react-i18next";
import {Dispatch, SetStateAction} from "react";

interface NavbarProps {
    setNavHide: Dispatch<SetStateAction<boolean>>;
}

export function Navbar({setNavHide}: NavbarProps) {
    const {t} = useTranslation();

    return (
        <Flex p={4} alignItems="center" pb={1} maxH={'10vh'}>
            <Tooltip label={t('btn.close')}>
                <IconButton icon={<Menu/>} aria-label="Menu" onClick={() => setNavHide(old => !old)}/>
            </Tooltip>
            <Text fontSize="xl" fontWeight="bold" ml={4}>WebCrawler</Text>
            <Spacer/>
            <PlayCrawler/>

            <ToggleTheme/>

            <ButtonGroup aria-label={t('btn.hideToTray')} mx={1}
                         onClick={() => ConfigService.HideWindow()}>
                <Tooltip label={t('btn.hideToTray')}>
                    <EyeOff/>
                </Tooltip>
            </ButtonGroup>

        </Flex>
    );
}