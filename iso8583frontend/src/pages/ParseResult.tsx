import {
    Accordion,
    ActionIcon,
    Button, Center,
    Container,
    createStyles, Flex,
    Group, JsonInput,
    LoadingOverlay,
    MultiSelect,
    ScrollArea,
    Table,
    Tabs,
    TextInput,
    Title
} from '@mantine/core';
import type {IParseResult, Search} from '../types.ts';
import {useEffect, useState} from 'react';
import {IconMessageCircle, IconPlus, IconSearch, IconTrash} from '@tabler/icons-react';
import {useHotkeys} from '@mantine/hooks';
import {GetResponseMessage} from '../types.ts';


const useStyles = createStyles((theme) => ({
    parseContainer: {
        ...theme.fn.focusStyles(),
        fontSize: theme.fontSizes.md,
        color: theme.colorScheme === 'dark' ? theme.colors.dark[1] : theme.colors.gray[7],
        height: 45,
        width: '100%',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: theme.spacing.md,
    }
}))


type ParseResultProps = {
    name: string,
}

export default function ParseResult({name}: ParseResultProps) {
    const {classes} = useStyles()
    const [parse, setParse] = useState<IParseResult>()
    const [selectedTab, setSelectedTab] = useState<string | null>('search')
    const [newFieldKey, setNewFieldKey] = useState<string>('')
    const [searchNumber, setSearchNumber] = useState<number>(0)

    useHotkeys([['ctrl+k', () => setSelectedTab(selectedTab === 'search' ? 'results' : 'search')]])

    useEffect(() => {
        fetch(`${import.meta.env.DEV ? 'http://127.0.0.1:8000' : ''}/api/v1/parse/${name}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(search)
        })
            .then(response => response.json())
            .then(data => {
                setParse(data)
            })
    }, [searchNumber])

    const [search, setSearch] = useState<Search>({
        fields: {
            '037': '',
        }
    })

    const HandleSearch = () => {
        console.log(search)
        setSelectedTab('results')
        setSearchNumber(searchNumber + 1)
    }

    return (
        <>
            {
                parse ?
                    <>
                        <Container className={classes.parseContainer}>
                            <Title order={2} my="md">{parse.name}</Title>
                            <Title order={2} my="md">{new Date(parse.date).toDateString()}</Title>
                        </Container>
                        <Tabs value={selectedTab} onTabChange={setSelectedTab}>
                            <Tabs.List>
                                <Tabs.Tab value="search" icon={<IconSearch size="0.8rem"/>}>Search</Tabs.Tab>
                                <Tabs.Tab value="results" icon={<IconMessageCircle size="0.8rem"/>}>Results</Tabs.Tab>
                            </Tabs.List>

                            <Tabs.Panel value="search" pt="xs">
                                <ScrollArea>
                                    <Container p="md">
                                        <Title order={2} my="md">MTI</Title>
                                        <Group p="md" grow>
                                            <TextInput
                                                label="Version"
                                                placeholder="0"
                                                value={search.mtiVersion ? search.mtiVersion : ''}
                                                onChange={(event) =>
                                                    setSearch({...search, mtiVersion: event.currentTarget.value})}
                                            />
                                            <TextInput
                                                label="Class"
                                                placeholder="2"
                                                value={search.mtiClass ? search.mtiClass : ''}
                                                onChange={(event) =>
                                                    setSearch({...search, mtiClass: event.currentTarget.value})}
                                            />
                                            <TextInput
                                                label="Function"
                                                placeholder="1"
                                                value={search.mtiFunction ? search.mtiFunction : ''}
                                                onChange={(event) =>
                                                    setSearch({...search, mtiFunction: event.currentTarget.value})}
                                            />
                                            <TextInput
                                                label="Origin"
                                                placeholder="0"
                                                value={search.mtiOrigin ? search.mtiOrigin : ''}
                                                onChange={(event) =>
                                                    setSearch({...search, mtiOrigin: event.currentTarget.value})}
                                            />
                                        </Group>
                                    </Container>
                                    <Container p="md">
                                        <Title order={2} my="md">Bitmap & LogFile</Title>
                                        <Group p="md" grow>
                                            <TextInput
                                                label="BitMap"
                                                placeholder="F02000140A0088200000000004000002"
                                                value={search.bitmap ? search.bitmap : ''}
                                                onChange={(event) =>
                                                    setSearch({
                                                        ...search,
                                                        bitmap: event.currentTarget.value.replace(/ /g, '')
                                                    })}
                                            />
                                            <MultiSelect
                                                data={parse.logFiles?.map((logFile) => ({value: logFile, label: logFile})) || []}
                                                label="Log files to search"
                                                placeholder="Select log files to search"
                                                searchable
                                                nothingFound="No log files found"
                                                value={search.logFiles ? search.logFiles : []}
                                                onChange={(value) => setSearch({...search, logFiles: value})}
                                            />
                                        </Group>
                                    </Container>
                                    <Container>
                                        <Flex justify="space-between" align="center">
                                            <Title order={2} my="md">Fields</Title>
                                            <Flex gap="xs">
                                                <TextInput
                                                    placeholder="Field num"
                                                    value={newFieldKey}
                                                    onChange={(event) => setNewFieldKey(event.currentTarget.value)}
                                                />
                                                <Button
                                                    variant="light"
                                                    rightIcon={<IconPlus />}
                                                    onClick={() => {
                                                        if (newFieldKey === '') return
                                                        setSearch({
                                                            ...search,
                                                            fields: {
                                                                ...search.fields,
                                                                [newFieldKey]: ''
                                                            }
                                                        })
                                                        setNewFieldKey('')
                                                    }}
                                                >
                                                    Add
                                                </Button>
                                            </Flex>
                                        </Flex>
                                        <Table
                                            highlightOnHover
                                            striped
                                        >
                                            <thead>
                                                <tr>
                                                    <th>Field Num</th>
                                                    <th>Field value</th>
                                                    <th></th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {
                                                    Object.entries(search.fields as { [key: string]: string }).map(([key, value]) => (
                                                        <tr key={key}>
                                                            <td>{key}</td>
                                                            <td>
                                                                <TextInput
                                                                    placeholder="Field value"
                                                                    value={value}
                                                                    onChange={(event) =>
                                                                        setSearch({
                                                                            ...search,
                                                                            fields: {
                                                                                ...search.fields,
                                                                                [key]: event.currentTarget.value
                                                                            }
                                                                        })}
                                                                />
                                                            </td>
                                                            <td
                                                                onClick={() => setSearch({
                                                                    ...search,
                                                                    fields: Object.fromEntries(
                                                                        Object.entries(search.fields as { [key: string]: string }).filter(([k]) => k !== key)
                                                                    )
                                                                })
                                                                }
                                                            >
                                                                <ActionIcon>
                                                                    <IconTrash />
                                                                </ActionIcon>
                                                            </td>
                                                        </tr>
                                                    ))
                                                }
                                            </tbody>
                                        </Table>
                                        <Center>
                                            <Button
                                                mt="xl"
                                                variant="light"
                                                rightIcon={<IconSearch />}
                                                onClick={HandleSearch}
                                            >
                                                Search
                                            </Button>
                                        </Center>
                                    </Container>
                                </ScrollArea>
                            </Tabs.Panel>
                            <Tabs.Panel value="results" pt="xs">
                                <Accordion variant="separated">
                                    {
                                        parse.messages.map((message, index) => (
                                            message.fields['037']?.value &&
                                                <Accordion.Item value={message.fields['037']?.value} key={index}>
                                                    <Accordion.Control>
                                                        037: {message.fields['037']?.value} <br/>
                                                        MTI: {message.mti.version}{message.mti.class}{message.mti.function}{message.mti.origin} <br/>
                                                        Response: {message.fields['039']?.value} ({GetResponseMessage(message.fields['039']?.value)}) <br/>
                                                    </Accordion.Control>
                                                    <Accordion.Panel>
                                                        <JsonInput
                                                            value={JSON.stringify(message, null, 2)}
                                                            autosize
                                                        />
                                                    </Accordion.Panel>
                                                </Accordion.Item>
                                        ))
                                    }
                                </Accordion>
                            </Tabs.Panel>
                        </Tabs>
                    </>
                    : <LoadingOverlay visible/>
            }
        </>
    )
}
