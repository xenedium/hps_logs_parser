import {Button, Container, Divider, Flex, Loader, Modal, SimpleGrid, Text, TextInput, Title} from '@mantine/core';
import {FloatingLabelInput} from '../components/FloatingLabelInput.tsx';
import {useState} from 'react';
import {GradientSegmentedControl} from '../components/GradientSegmentedControl.tsx';
import {DropzoneButton} from '../components/DropzoneButton.tsx';
import {FileWithPath} from '@mantine/dropzone';

export default function NewParse() {

    const HandleCreateParse = async () => {
        setLoading(true)
        if (newParse.type === 'ssh') {
            const response = await fetch(`${import.meta.env.DEV ? 'http://127.0.0.1:8000' : ''}/api/v1/ssh`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    host: newParse.ssh.host,
                    port: newParse.ssh.port,
                    user: newParse.ssh.username,
                    password: newParse.ssh.password,
                    absoluteDir: newParse.ssh.path,
                    parseRequestName: newParse.name,
                })
            })
            const data = await response.json()
            setLoading(false)
            if (data.error) return alert(data.error)
            alert('Parse request created successfully')
            window.location.reload()
        }
        if (newParse.type === 'upload') {
            const formData = new FormData()
            newParse.upload.files.forEach((file) => formData.append('files', file))
            formData.append('parseRequestName', newParse.name)

            const response = await fetch(`${import.meta.env.DEV ? 'http://127.0.0.1:8000' : ''}/api/v1/upload`, {
                method: 'POST',
                body: formData
            })

            const data = await response.json()
            setLoading(false)
            if (data.error) return alert(data.error)
            alert('Parse request created successfully')
            window.location.reload()
        }
    }


    const [newParse, setNewParse] = useState<{
        name: string
        type: 'ssh' | 'upload'
        ssh: {
            host: string
            port: number
            username: string
            password: string
            path: string
        },
        upload: {
            files: FileWithPath[]
        }
    }
    >({
        name: '',
        type: 'ssh',
        ssh: {
            host: '',
            port: 22,
            username: '',
            password: '',
            path: '',
        },
        upload: {
            files: [],
        }
    })

    const [loading, setLoading] = useState<boolean>(false)

    return (
        <div>
            <Modal centered opened={loading} onClose={() => setLoading(false)} withCloseButton={false}>
                <Flex align="center" gap="xl">
                    <Loader size="xl"/>
                    <div>
                        <Text>Creating parse request...</Text>
                        <Text c="dimmed">This can take seconds up to minutes...</Text>
                    </div>
                </Flex>
            </Modal>
            <Title p="md">New Parse Request</Title>
            <Container p="md">
                <Container>
                    <Title order={2} my="md">General Info</Title>
                    <SimpleGrid cols={1}>
                        <GradientSegmentedControl
                            value={newParse.type}
                            onChange={(value) => setNewParse({...newParse, type: value as 'ssh' | 'upload'})}
                            data={['ssh', 'upload']}
                        />
                        <FloatingLabelInput
                            label="Parse name"
                            placeholder="ssh-parse-01/01/2023"
                            value={newParse.name}
                            onChange={(event) => setNewParse({...newParse, name: event.currentTarget.value})}
                        />
                    </SimpleGrid>
                </Container>
                <Divider my="md"/>
                <Container>
                    <Title order={2} my="md">Details</Title>
                    {
                        newParse.type === 'ssh' ?
                            <SimpleGrid cols={2}>
                                <TextInput
                                    label="SSH Host"
                                    placeholder="20.20.20.20 or example.com"
                                    required
                                    value={newParse.ssh.host}
                                    onChange={(event) => setNewParse({
                                        ...newParse,
                                        ssh: {...newParse.ssh, host: event.currentTarget.value}
                                    })}
                                />
                                <TextInput
                                    label="SSH Port"
                                    placeholder="22"
                                    required
                                    value={newParse.ssh.port}
                                    onChange={(event) => setNewParse({
                                        ...newParse,
                                        ssh: {...newParse.ssh, port: parseInt(event.currentTarget.value)}
                                    })}
                                />
                                <TextInput
                                    label="SSH Username"
                                    placeholder="root"
                                    required
                                    value={newParse.ssh.username}
                                    onChange={(event) => setNewParse({
                                        ...newParse,
                                        ssh: {...newParse.ssh, username: event.currentTarget.value}
                                    })}
                                />
                                <TextInput
                                    label="SSH Password"
                                    placeholder="password"
                                    required
                                    value={newParse.ssh.password}
                                    onChange={(event) => setNewParse({
                                        ...newParse,
                                        ssh: {...newParse.ssh, password: event.currentTarget.value}
                                    })}
                                />
                                <TextInput
                                    label="Absolute Path"
                                    placeholder="/var/logs"
                                    required
                                    value={newParse.ssh.path}
                                    onChange={(event) => setNewParse({
                                        ...newParse,
                                        ssh: {...newParse.ssh, path: event.currentTarget.value}
                                    })}
                                />
                            </SimpleGrid>
                            :
                            <DropzoneButton
                                files={newParse.upload.files}
                                onDrop={(files) => setNewParse({...newParse, upload: {files: files}})}
                            />
                    }
                </Container>
                <Container mt="xl">
                    <Button
                        variant="light"
                        radius="xl"
                        fullWidth
                        onClick={HandleCreateParse}
                    >
                        Create Parse Request
                    </Button>
                </Container>
            </Container>
        </div>
    )
}