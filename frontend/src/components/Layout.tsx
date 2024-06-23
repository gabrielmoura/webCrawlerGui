import {Box, Divider, Flex} from "@chakra-ui/react";
import {ReactNode, useState} from "react";
import {SideBar, SidebarStyles} from "./Sidebar.tsx";
import {Navbar} from "./Navbar.tsx";

interface FRootLayoutProps {
    children?: ReactNode;
}

export function FRootLayout({children}: FRootLayoutProps) {
    const [navHide, setNavHide] = useState(false)

    return (
        <>
            <Flex minH="100vh" direction="column">
                <Flex flex="1" overflow="hidden">
                    {!navHide ? <SideBar/> : null}
                    <Box flex="1" bg="gray.suave" ml={!navHide ? SidebarStyles.sideBox.width : "0"} overflow="auto">
                        <Navbar setNavHide={setNavHide}/>
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