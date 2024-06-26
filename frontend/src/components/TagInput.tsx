import {Dispatch, SetStateAction, useState} from "react";
import {Button, Flex, Input, InputGroup, InputRightElement, Text, Tooltip,} from "@chakra-ui/react";
import {Ban, Plus} from "lucide-react";
import {useTranslation} from "react-i18next";
import {onEnter} from "../util/helper.ts";

interface TagInputProps {
    tags: string[];
    setTags: Dispatch<SetStateAction<string[]>>;
    placeholder: string;
    regex?: RegExp;
}

export const TagInput = ({
                             tags,
                             setTags,
                             placeholder,
                             regex,
                         }: TagInputProps) => {
    const {t} = useTranslation();
    const [isInputVisible, setInputVisible] = useState(false);
    const [err, setErr] = useState(false);
    const [newTag, setNewTag] = useState("");

    const handleAddTag = () => {
        if (newTag.trim() !== "") {
            if (regex && !regex.test(newTag.trim())) {
                setErr(true);
            } else {
                setTags([...tags, newTag.trim()]);
                cancelInput();
            }
        }
    };

    const showInput = () => {
        setInputVisible(true);
    };
    const flushData = () => {
        setTags([]);
    };
    const cancelInput = () => {
        setNewTag("");
        setInputVisible(false);
        setErr(false);
    };

    return (
        <Flex align="center">
            <Text
                mr={2}
                borderWidth={"1px"}
                minW={"22rem"}
                p={2} // Added padding for spacing
                id={`tagInput-result`}
            >
                {tags.length > 0 ? (
                    tags.map((tag, index) => (
                        <span key={index}>
              {tag}
                            {index !== tags.length - 1 && ", "}
            </span>
                    ))
                ) : (
                    <span>&nbsp;</span>
                )}
            </Text>
            {!isInputVisible ? (
                <Flex gap={"0.2rem"}>
                    <Button size="sm" onClick={showInput}>
                        <Tooltip label={t('btn.include')}>
                            <Plus/>
                        </Tooltip>
                    </Button>
                    <Button size="sm" onClick={flushData}>
                        <Tooltip label={t('btn.removeAll')}>
                            <Ban/>
                        </Tooltip>
                    </Button>
                </Flex>
            ) : (
                <InputGroup size="md">
                    <Input
                        id={`tagInput-input`}
                        pr="4.5rem"
                        isInvalid={err}
                        placeholder={placeholder}
                        value={newTag}
                        onKeyDown={e => onEnter(e, handleAddTag)}
                        onChange={(e) => setNewTag(e.target.value)}
                    />
                    <InputRightElement width="8rem">
                        <Flex gap={"0.2rem"}>
                            <Tooltip label="Add">
                                <Button
                                    h="1.75rem"
                                    size="sm"
                                    ml={0}
                                    onClick={handleAddTag}
                                    id={`tagInput-btnAcr`}
                                >
                                    Add
                                </Button>
                            </Tooltip>
                            <Tooltip label={t('btn.cancel')}>
                                <Button
                                    h="1.75rem"
                                    size="sm"
                                    ml={0}
                                    onClick={cancelInput}
                                    id={`tagInput-btnCan`}
                                >
                                    <Ban/>
                                </Button>
                            </Tooltip>
                        </Flex>
                    </InputRightElement>
                </InputGroup>
            )}
        </Flex>
    );
};
