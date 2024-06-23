import React, {useState} from 'react';
import {Box, Collapse, IconButton, Text, VStack} from '@chakra-ui/react';
import {ChevronDown, ChevronRight} from "lucide-react";

interface TreeNode {
    title: string;
    description?: string;
    children?: TreeNode[];
}

interface TreeProps {
    nodes: TreeNode[];
}

const Tree: React.FC<TreeProps> = ({nodes}) => {
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

    const toggle = () => setIsOpen(!isOpen);

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
                    <Text fontWeight="bold">{node.title}</Text>
                    {node.description && (
                        <Text fontSize="sm" color="gray.600">
                            {node.description}
                        </Text>
                    )}
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

export default Tree;
