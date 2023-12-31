import {
    ActionIcon,
    Button,
    createStyles,
    Divider,
    Flex,
    getStylesRef,
    Navbar,
    ScrollArea,
    UnstyledButton
} from '@mantine/core';
import {SegmentedToggle} from './SegmentedToggle.tsx';
import {IconPlus, IconSearch, IconTrash} from '@tabler/icons-react';


type SidenavProps = {
    parses: string[]
    selection: string | 'new-parse'
    setSelection: (id: string | 'new-parse') => void
}

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
        borderRadius: theme.radius.md,
        cursor: 'pointer',
        '&:hover': {
            backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[2],
            color: theme.colorScheme === 'dark' ? theme.white : theme.black,

            [`& .${getStylesRef('icon')}`]: {
                color: theme.colorScheme === 'dark' ? theme.white : theme.black,
            },
        }
    },
    parseContainerActive: {
        '&, &:hover': {
            backgroundColor: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).background,
            color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
            [`& .${getStylesRef('icon')}`]: {
                color: theme.fn.variant({ variant: 'light', color: theme.primaryColor }).color,
            },
        },
    }
}))

export function Sidenav({parses, selection, setSelection}: SidenavProps) {

    const {classes, cx} = useStyles()

    const HandleDelete = (id: string) => {
        fetch(`${import.meta.env.DEV ? 'http://127.0.0.1:8000' : ''}/api/v1/keys/${id}`, {
            method: 'DELETE'
        })
            .then(r => r.json())
            .then(data => {
                if (data.status === 'ok') {
                    window.location.reload()
                    return
                }
                alert('Something went wrong')
            })
            .catch(() => alert('Something went wrong'))
    }

    return (
        <Navbar width={{base: 300}} p="md">
            <Navbar.Section>
                <Button
                    variant="light"
                    radius="xl"
                    fullWidth
                    rightIcon={<IconPlus size={18}/>}
                    onClick={() => setSelection('new-parse')}
                >
                    New Parse
                </Button>
            </Navbar.Section>

            <Divider my="md"/>

            <Navbar.Section grow>
                <ScrollArea
                    offsetScrollbars
                    type="hover"
                    h="calc(90vh - 100px)"
                >
                    <Flex
                        direction="column"
                        gap="xs"
                    >
                        {
                            parses.map((parse) => {
                                return (
                                    <UnstyledButton
                                        key={parse}
                                        onClick={() => selection !== parse && setSelection(parse)}
                                        className={cx(classes.parseContainer, {[classes.parseContainerActive]: selection === parse})}
                                    >
                                        <Flex
                                            align="center"
                                            gap="md"
                                        >
                                            <IconSearch size={18}/>
                                            {parse}
                                        </Flex>
                                        {
                                            selection === parse &&
                                            <ActionIcon
                                                variant="outline"
                                                onClick={() => HandleDelete(parse)}
                                            >
                                                <IconTrash size={18}/>
                                            </ActionIcon>
                                        }

                                    </UnstyledButton>
                                )
                            })
                        }
                    </Flex>
                </ScrollArea>
            </Navbar.Section>

            <Divider my="md"/>

            <Navbar.Section>
                <SegmentedToggle/>
            </Navbar.Section>
        </Navbar>
    )
}