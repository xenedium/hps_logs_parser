import {useRef} from 'react';
import {Button, createStyles, Group, rem, SimpleGrid, Text} from '@mantine/core';
import {Dropzone, FileWithPath} from '@mantine/dropzone';
import {IconCloudUpload, IconDownload, IconX} from '@tabler/icons-react';

const useStyles = createStyles((theme) => ({
    wrapper: {
        position: 'relative',
        marginBottom: rem(60),
    },

    dropzone: {
        borderWidth: rem(1),
        paddingBottom: rem(50),
    },

    icon: {
        color: theme.colorScheme === 'dark' ? theme.colors.dark[3] : theme.colors.gray[4],
    },

    control: {
        position: 'absolute',
        width: rem(250),
        left: `calc(50% - ${rem(125)})`,
        bottom: rem(-20),
    },
}));

type DropzoneButtonProps = {
    files: FileWithPath[];
    onDrop: (files: FileWithPath[]) => void;
};

export function DropzoneButton({files, onDrop}: DropzoneButtonProps) {
    const {classes, theme} = useStyles();
    const openRef = useRef<() => void>(null);

    return (
        <>
            <div className={classes.wrapper}>
                <Dropzone
                    openRef={openRef}
                    onDrop={onDrop}
                    className={classes.dropzone}
                    radius="md"
                    maxSize={30 * 1024 ** 2}
                >
                    <div style={{pointerEvents: 'none'}}>
                        <Group position="center">
                            <Dropzone.Accept>
                                <IconDownload
                                    size={rem(50)}
                                    color={theme.colors[theme.primaryColor][6]}
                                    stroke={1.5}
                                />
                            </Dropzone.Accept>
                            <Dropzone.Reject>
                                <IconX size={rem(50)} color={theme.colors.red[6]} stroke={1.5}/>
                            </Dropzone.Reject>
                            <Dropzone.Idle>
                                <IconCloudUpload
                                    size={rem(50)}
                                    color={theme.colorScheme === 'dark' ? theme.colors.dark[0] : theme.black}
                                    stroke={1.5}
                                />
                            </Dropzone.Idle>
                        </Group>

                        <Text ta="center" fw={700} fz="lg" mt="xl">
                            <Dropzone.Accept>Drop files here</Dropzone.Accept>
                            <Dropzone.Idle>Upload log files</Dropzone.Idle>
                        </Text>
                        <Text ta="center" fz="sm" mt="xs" c="dimmed">
                            Drag&apos;n&apos;drop files here to upload. There's no limit on the number of files you can
                            upload.
                        </Text>
                    </div>
                </Dropzone>

                <Button className={classes.control} size="md" radius="xl" onClick={() => openRef.current?.()}>
                    Select files
                </Button>
            </div>
            {
                files.length > 0 &&
                <>
                    <Text>Selected files: </Text>
                    <SimpleGrid cols={2} spacing="md">
                        {files.map((file) => (
                            <div key={file.path}>
                                <Text lineClamp={2}>{file.name}</Text>
                            </div>
                        ))}
                    </SimpleGrid>
                </>
            }
        </>
    );
}