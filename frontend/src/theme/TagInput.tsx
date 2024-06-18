import { Dispatch, SetStateAction, useState, KeyboardEvent } from "react";
import {
  Text,
  Flex,
  Button,
  Input,
  InputGroup,
  InputRightElement,
  Tooltip,
} from "@chakra-ui/react";
import { Ban, Plus } from "lucide-react";

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

  const handleInputKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleAddTag();
    }
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
            <Tooltip label="Incluir">
              <Plus />
            </Tooltip>
          </Button>
          <Button size="sm" onClick={flushData}>
            <Tooltip label="Remover todos">
              <Ban />
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
            onKeyDown={handleInputKeyDown}
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
              <Tooltip label="Cencelar">
                <Button
                  h="1.75rem"
                  size="sm"
                  ml={0}
                  onClick={cancelInput}
                  id={`tagInput-btnCan`}
                >
                  <Ban />
                </Button>
              </Tooltip>
            </Flex>
          </InputRightElement>
        </InputGroup>
      )}
    </Flex>
  );
};
