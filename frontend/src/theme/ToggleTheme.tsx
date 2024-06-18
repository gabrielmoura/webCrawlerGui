import {ButtonGroup, Tooltip, useColorMode} from "@chakra-ui/react";
import {Moon, Sun} from "lucide-react";
import useAppStore from "../store/appStore.ts";
import {useTranslation} from "react-i18next";

export function ToggleTheme() {
    const {t} = useTranslation();
    const setWTheme = useAppStore(s => s.setWindowTheme)
    const {colorMode, toggleColorMode} = useColorMode();


    function toggleCTheme() {
        toggleColorMode()
        setWTheme(colorMode == 'dark')
    }

    return <ButtonGroup aria-label="Theme" mx={1} onClick={() => toggleCTheme()}>
        <Tooltip label={t('theme')}>
            {colorMode == 'dark' ? <Sun color={'#FFD700'}/> : <Moon color={'#4169E1'}/>}
        </Tooltip>
    </ButtonGroup>
}