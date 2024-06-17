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
    VStack
} from "@chakra-ui/react";
import {Link} from "@tanstack/react-router";
import imageLogo from "../assets/images/logo-universal.png";
import React, {useState} from "react";
import {Bolt, EyeOff, Menu, Rows4, Search} from "lucide-react";
import {PlayCrawler} from "./PlayCrawler.tsx";
import {ToggleTheme} from "./ToggleTheme.tsx";
import {Hide} from "../../wailsjs/runtime";

interface FRootLayoutProps {
    children?: React.ReactNode;
}

export function FRootLayout({children}: FRootLayoutProps) {
    const [navHide, setNavHide] = useState(false)
    const iconStyle = {
        marginTop: '2px',
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
                        <Flex p={4} alignItems="center" pb={1} maxH={'10vh'}>
                            <Tooltip label='Fechar SideBar'>
                                <IconButton icon={<Menu/>} aria-label="Menu" onClick={() => setNavHide(old => !old)}/>
                            </Tooltip>
                            <Text fontSize="xl" fontWeight="bold" ml={4}>Constellation</Text>
                            <Spacer/>
                            <PlayCrawler/>

                            <ToggleTheme/>

                            {/*<ButtonGroup aria-label="Settings" mx={1}>*/}
                            {/*    <Tooltip label='Settings'>*/}
                            {/*        <Bolt/>*/}
                            {/*    </Tooltip>*/}
                            {/*</ButtonGroup>*/}

                            {/*<ButtonGroup aria-label="Logout" mx={1}>*/}
                            {/*    <Tooltip label='Logout'>*/}
                            {/*        <LogOut/>*/}
                            {/*    </Tooltip>*/}
                            {/*</ButtonGroup>*/}

                            <ButtonGroup aria-label="Logout" mx={1} onClick={() => Hide()}>
                                <Tooltip label='Ocultar para a bandeja'>
                                    <EyeOff/>
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