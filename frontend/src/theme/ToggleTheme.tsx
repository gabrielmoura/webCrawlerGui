import {ButtonGroup, Tooltip, useColorMode} from "@chakra-ui/react";
import {Moon, Sun} from "lucide-react";
import useAppStore from "../store/appStore.ts";

export function ToggleTheme() {
    const setWTheme = useAppStore(s => s.setWindowTheme)
    const {colorMode, toggleColorMode} = useColorMode();

    function toggleCTheme() {
        toggleColorMode()
        setWTheme(colorMode == 'dark')
    }

    return <ButtonGroup aria-label="Theme" mx={1} onClick={() => toggleCTheme()}>
        <Tooltip label='Theme'>
            {colorMode == 'dark' ? <Sun color={'#FFD700'}/> : <Moon color={'#4169E1'}/>}
        </Tooltip>
    </ButtonGroup>
}