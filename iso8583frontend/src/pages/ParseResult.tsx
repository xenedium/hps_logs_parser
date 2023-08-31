import {
    Container,
    createStyles,
    Group,
    LoadingOverlay,
    MultiSelect,
    ScrollArea,
    Tabs,
    TextInput,
    Title
} from '@mantine/core';
import type {IParseResult, Search} from '../types.ts';
import {useEffect, useState} from 'react';
import {IconMessageCircle, IconSearch} from '@tabler/icons-react';
import {useHotkeys} from '@mantine/hooks';


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
    id: number,
}

export default function ParseResult({id}: ParseResultProps) {
    const {classes} = useStyles()
    const [parse, setParse] = useState<IParseResult>()
    const [selectedTab, setSelectedTab] = useState<string | null>('search')

    useHotkeys([['ctrl+k', () => setSelectedTab(selectedTab === 'search' ? 'results' : 'search')]])

    useEffect(() => {
        setParse({
            id: 1,
            name: 'parse-result-1',
            date: new Date(),
            type: 'ssh',
            status: 'done',
            messages: [
                {
                    'fields': {
                        '102': {
                            'length': '0',
                            'value': '2000039098',
                            'raw': ''
                        },
                        '002': {
                            'length': '0',
                            'value': '4301400012814919',
                            'raw': ''
                        },
                        '004': {
                            'length': '0',
                            'value': '000000000000',
                            'raw': ''
                        },
                        '059': {
                            'length': '0',
                            'value': '0000000326',
                            'raw': ''
                        },
                        '127.025': {
                            'length': '0',
                            'value': '04000000000000000004',
                            'raw': ''
                        },
                        '003': {
                            'length': '0',
                            'value': '310000',
                            'raw': ''
                        },
                        '011': {
                            'length': '0',
                            'value': '000225',
                            'raw': ''
                        },
                        '037': {
                            'length': '0',
                            'value': '000000599724',
                            'raw': ''
                        },
                        '028': {
                            'length': '0',
                            'value': 'C00000000',
                            'raw': ''
                        },
                        '039': {
                            'length': '0',
                            'value': '96',
                            'raw': ''
                        },
                        '049': {
                            'length': '0',
                            'value': '404',
                            'raw': ''
                        },
                        '053': {
                            'length': '0',
                            'value': '303030303030303033363C30303030303C30303030303032353032353F303030303',
                            'raw': ''
                        },
                        '030': {
                            'length': '0',
                            'value': 'C00000000',
                            'raw': ''
                        }
                    },
                    'mti': {
                        'version': 0,
                        'class': 2,
                        'function': 1,
                        'origin': 0
                    },
                    'bitmap': 'F02000140A0088200000000004000002',
                    'raw': '',
                    'logFileName': '/tmp/3070292616/POSTILION.TRC000',
                    'lineNumber': '3948'
                },
                {
                    'fields': {
                        '102': {
                            'length': '0',
                            'value': '7180003668',
                            'raw': ''
                        },
                        '004': {
                            'length': '0',
                            'value': '000000000000',
                            'raw': ''
                        },
                        '037': {
                            'length': '0',
                            'value': '000000599734',
                            'raw': ''
                        },
                        '053': {
                            'length': '0',
                            'value': '303030303030303033313C30303030303C30303030303031383031383F303030303',
                            'raw': ''
                        },
                        '059': {
                            'length': '0',
                            'value': '0000000331',
                            'raw': ''
                        },
                        '028': {
                            'length': '0',
                            'value': 'C00000000',
                            'raw': ''
                        },
                        '030': {
                            'length': '0',
                            'value': 'C00000000',
                            'raw': ''
                        },
                        '039': {
                            'length': '0',
                            'value': '96',
                            'raw': ''
                        },
                        '127.025': {
                            'length': '0',
                            'value': '04000000000000000004',
                            'raw': ''
                        },
                        '002': {
                            'length': '0',
                            'value': '4301400038595500',
                            'raw': ''
                        },
                        '003': {
                            'length': '0',
                            'value': '310000',
                            'raw': ''
                        },
                        '011': {
                            'length': '0',
                            'value': '000148',
                            'raw': ''
                        },
                        '049': {
                            'length': '0',
                            'value': '404',
                            'raw': ''
                        }
                    },
                    'mti': {
                        'version': 0,
                        'class': 2,
                        'function': 1,
                        'origin': 0
                    },
                    'bitmap': 'F02000140A0088200000000004000002',
                    'raw': '',
                    'logFileName': '/tmp/3070292616/POSTILION.TRCdd000',
                    'lineNumber': '13581'
                }
            ]
        })
    }, [id])

    const [search, setSearch] = useState<Search>({})


    return (
        <>
            {
                parse ?
                    <>
                        <Container className={classes.parseContainer}>
                            <Title order={2} my="md">{parse.name}</Title>
                            <Title order={2} my="md">{parse.date.toDateString()}</Title>
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
                                                data={parse.messages.map((message) => message.logFileName)}
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
                                        <Title order={2} my="md">Fields</Title>
                                        {/* TODO: make a table with fields */}
                                    </Container>
                                </ScrollArea>
                            </Tabs.Panel>

                            <Tabs.Panel value="results" pt="xs">
                                Messages tab content
                            </Tabs.Panel>

                        </Tabs>
                    </>
                    : <LoadingOverlay visible/>
            }
        </>
    )
}