import React, {useState} from 'react';
import {Box, Center, Collapse, IconButton, Text, Tooltip, useToast, VStack} from '@chakra-ui/react';
import {ChevronDown, ChevronRight, CircleX} from "lucide-react";
import {useTranslation} from "react-i18next";
import {useMutation} from "@tanstack/react-query";
import {QueueService} from "../services/queue.ts";

export interface TreeNode {
    title: string;
    description?: string;
    url: string;
    children?: TreeNode[];

}

interface TreeProps {
    nodes: TreeNode[];
}

const TreeLink: React.FC<TreeProps> = ({nodes}) => {
    const renderTree = (nodes: TreeNode[]) => {
        return nodes.map((node, index) => <TreeNodeComponent key={index} node={node}/>);
    };

    return <Box>{renderTree(nodes)}</Box>;
};

interface TreeNodeComponentProps {
    node: TreeNode;
}

const TreeNodeComponent: React.FC<TreeNodeComponentProps> = ({node}) => {
    const [isOpen, setIsOpen] = useState(false);
    const hasChildren = node.children && node.children.length > 0;
    const {t} = useTranslation();
    const toast = useToast()

    const toggle = () => setIsOpen(!isOpen);

    function handleDeleteFromQueue(url: string) {
        if (url) {
            mutateDelete.mutate(url)
        }
    }

    const mutateDelete = useMutation({
        mutationKey: ['queue', 'delete'],
        mutationFn: async (url: string) => QueueService.removeFromQueue(url),
        onSuccess: async (msg) => {
            toast({
                title: t(`msg.${msg}`),
                status: 'success',
                duration: 9000,
                isClosable: true,
            })
        }, onError: (msg) => {
            toast({
                title: t(`msg.${msg}`),
                status: 'error',
                duration: 9000,
                isClosable: true,
            })
        }
    })

    return (
        <VStack align="start" pl={4} mt={2} spacing={1} borderLeft="2px solid" borderColor="gray.200">
            <Box display="flex" alignItems="center">
                {hasChildren && (
                    <IconButton
                        icon={isOpen ? <ChevronDown/> : <ChevronRight/>}
                        size="sm"
                        onClick={toggle}
                        aria-label={isOpen ? 'Collapse' : 'Expand'}
                        variant="ghost"
                    />
                )}
                <Box>
                    <Center>
                        <Box>
                            <Text fontWeight="bold" fontSize='2xl'>{node.title??node.url}</Text>
                            {node.description && (
                                <Text fontSize="sm" color="gray.600">
                                    {node.description}
                                </Text>
                            )}
                        </Box>
                        <Tooltip label={t('btn.delete')}>
                        <IconButton
                            onClick={() => handleDeleteFromQueue(node.url)}
                            aria-label={t('btn.delete')}
                            icon={<CircleX/>}
                            size='sm'
                        />
                        </Tooltip>
                    </Center>

                </Box>
            </Box>
            {hasChildren && (
                <Collapse in={isOpen} animateOpacity>
                    <Box pl={4}>{node.children && node.children.map((childNode, index) => <TreeNodeComponent key={index}
                                                                                                             node={childNode}/>)}</Box>
                </Collapse>
            )}
        </VStack>
    );
};

export default TreeLink;
