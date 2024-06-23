import {CSSProperties} from "react";
import {Box, Center, Divider, Image, Spacer, Text, VStack} from "@chakra-ui/react";
import imageLogo from "../assets/images/2-bg.png";
import {Link} from "@tanstack/react-router";
import {Bolt, Rows4, Search} from "lucide-react";
import {useTranslation} from "react-i18next";

export const SidebarStyles: Record<string, CSSProperties> = {
    sideBox: {
        width: '15.625rem',
        background: '#212528',
        color: 'white',
        padding: '1rem',
        position: 'fixed',
        height: '100vh',
        overflowY: 'auto'
    },
    iconStyle: {
        marginTop: '2px',
    }
}

export function SideBar() {
    const {t} = useTranslation();

    return (
        <Box style={SidebarStyles.sideBox}>
            <VStack align="start" spacing={4}>
                <Center w="100%">
                    <Image src={imageLogo} alt="Constellation Logo" boxSize="100px"/>
                </Center>
                <Text fontWeight="bold">{t('msg.welcome')}</Text>
                <Divider/>
                <VStack align="start" spacing={8}>
                    <Link to='/'>
                        <Center>
                            <Search style={SidebarStyles.iconStyle}/>
                            <Spacer w={2}/>
                            {t('nav.search')}
                        </Center>
                    </Link>
                    <Link to='/conf'>
                        <Center>
                            <Bolt style={SidebarStyles.iconStyle}/>
                            <Spacer w={2}/>
                            {t('nav.configs')}
                        </Center>
                    </Link>
                    <Link to='/queue'>
                        <Center>
                            <Rows4 style={SidebarStyles.iconStyle}/>
                            <Spacer w={2}/>
                            {t('nav.queue')}
                        </Center>
                    </Link>
                </VStack>
            </VStack>
        </Box>
    )
}