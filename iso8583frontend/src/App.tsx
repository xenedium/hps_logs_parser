import {AppShell} from '@mantine/core';

import {Sidenav} from './components/Sidenav.tsx';
import {useEffect, useState} from 'react';
import NewParse from './pages/NewParse.tsx';

const mockData = [
    {
        id: 1,
        name: 'parse-request-1',
        date: new Date(),
    },
    {
        id: 2,
        name: 'parse-request-2',
        date: new Date(),
    },
    {
        id: 3,
        name: 'parse-request-3',
        date: new Date(),
    },
    {
        id: 4,
        name: 'parse-request-4',
        date: new Date(),
    },
    {
        id: 5,
        name: 'parse-request-5',
        date: new Date(),
    },
    {
        id: 6,
        name: 'parse-request-6',
        date: new Date(),
    },
    {
        id: 7,
        name: 'parse-request-7',
        date: new Date(),
    },
    {
        id: 8,
        name: 'parse-request-8',
        date: new Date(),
    },
    {
        id: 9,
        name: 'parse-request-9',
        date: new Date(),
    },
    {
        id: 10,
        name: 'parse-request-10',
        date: new Date(),
    },
    {
        id: 11,
        name: 'parse-request-11',
        date: new Date(),
    },
    {
        id: 12,
        name: 'parse-request-12',
        date: new Date(),
    },
    {
        id: 13,
        name: 'parse-request-13',
        date: new Date(),
    },
    {
        id: 14,
        name: 'parse-request-14',
        date: new Date(),
    },
    {
        id: 15,
        name: 'parse-request-15',
        date: new Date(),
    },
    {
        id: 16,
        name: 'parse-request-16',
        date: new Date(),
    },
    {
        id: 17,
        name: 'parse-request-17',
        date: new Date(),
    },
    {
        id: 18,
        name: 'parse-request-18',
        date: new Date(),
    },
    {
        id: 19,
        name: 'parse-request-19',
        date: new Date(),
    },
    {
        id: 20,
        name: 'parse-request-20',
        date: new Date(),
    },
    {
        id: 21,
        name: 'parse-request-21',
        date: new Date(),
    },
    {
        id: 22,
        name: 'parse-request-22',
        date: new Date(),
    },
    {
        id: 23,
        name: 'parse-request-23',
        date: new Date(),
    },
    {
        id: 24,
        name: 'parse-request-24',
        date: new Date(),
    },
    {
        id: 25,
        name: 'parse-request-25',
        date: new Date(),
    }
]

// TODO: implement routing
function App() {
    const [parses, setParses] = useState<typeof mockData>([]);
    const [selection, setSelection] = useState<number>(1);

    useEffect(() => {
        // TODO: fetch parses from backend
        setParses(mockData)
    }, [])

    return (
        <AppShell
            padding="md"
            navbar={<Sidenav parses={parses} selection={selection} setSelection={(id: number) => setSelection(id)}/>}
        >
            <NewParse />
        </AppShell>
    )
}

export default App
