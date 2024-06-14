import {
    Box,
    ButtonGroup,
    Center,
    Divider,
    Flex,
    IconButton,
    Image,
    Spacer,
    Text,
    Tooltip,
    useColorMode,
    VStack
} from "@chakra-ui/react";
import {Link} from "@tanstack/react-router";
import imageLogo from "../assets/images/logo-universal.png";
import React, {useState} from "react";
import {Bolt, LogOut, Menu, Moon, Rows4, Search, Sun} from "lucide-react";
import useAppStore from "../store/appStore";

interface FRootLayoutProps {
    children?: React.ReactNode;
}

export function FRootLayout({children}: FRootLayoutProps) {
    const [navHide, setNavHide] = useState(false)
    const iconStyle = {
        marginTop: '2px',
    }
    const setWTheme = useAppStore(s => s.setWindowTheme)
    const {colorMode, toggleColorMode} = useColorMode();

    function toggleCTheme() {
        toggleColorMode()
        setWTheme(colorMode == 'dark')
    }

    return (
        <>
            <Flex minH="100vh" direction="column">
                <Flex flex="1" overflow="hidden">

                    {!navHide ? <Box w="250px" bg="#212528" color="white" p={4}>
                        <VStack align="start" spacing={4}>
                            <Center w="100%">
                                <Image src={imageLogo} alt="Constellation Logo" boxSize="100px"/>
                            </Center>
                            <Text fontWeight="bold">Bem vindo ao Buscador</Text>
                            <Divider/>
                            <VStack align="start" spacing={8}>
                                <Link to='/'>
                                    <Center>
                                        <Search style={iconStyle}/>
                                        <Spacer w={2}/>
                                        Buscar
                                    </Center>
                                </Link>
                                <Link to='/conf'>
                                    <Center>
                                        <Bolt style={iconStyle}/>
                                        <Spacer w={2}/>
                                        Configurações
                                    </Center>
                                </Link>
                                <Link to='/queue'>
                                    <Center>
                                        <Rows4 style={iconStyle}/>
                                        <Spacer w={2}/>
                                        Fila
                                    </Center>
                                </Link>
                            </VStack>
                        </VStack>
                    </Box> : null}


                    <Box flex="1" bg="gray.suave" overflow="auto">
                        <Flex p={4} alignItems="center" pb={1}>
                            <Tooltip label='Fechar SideBar'>
                                <IconButton icon={<Menu/>} aria-label="Menu" onClick={() => setNavHide(old => !old)}/>
                            </Tooltip>
                            <Text fontSize="xl" fontWeight="bold" ml={4}>Constellation</Text>
                            <Spacer/>

                            <ButtonGroup aria-label="Theme" mx={1} onClick={() => toggleCTheme()}>
                                <Tooltip label='Theme'>
                                    {colorMode == 'dark' ? <Sun/> : <Moon/>}
                                </Tooltip>
                            </ButtonGroup>

                            <ButtonGroup aria-label="Settings" mx={1}>
                                <Tooltip label='Settings'>
                                    <Bolt/>
                                </Tooltip>
                            </ButtonGroup>

                            <ButtonGroup aria-label="Logout" mx={1}>
                                <Tooltip label='Logout'>
                                    <LogOut/>
                                </Tooltip>
                            </ButtonGroup>

                        </Flex>
                        <Divider color='white'/>
                        <Box p={4}>
                            {children}
                        </Box>
                    </Box>
                </Flex>
            </Flex>
        </>
    );

}