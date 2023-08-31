import {AppShell} from '@mantine/core';

import {Sidenav} from './components/Sidenav.tsx';
import {useEffect, useState} from 'react';
import NewParse from './pages/NewParse.tsx';
import ParseResult from './pages/ParseResult.tsx';
import {IParseResult} from './types.ts';

const mockData: Omit<IParseResult, 'messages'>[] = [
    {
        id: 1,
        name: 'parse-result-1',
        date: new Date(),
        type: 'ssh',
        status: 'done',
    }
]

function App() {
    const [parses, setParses] = useState<Omit<IParseResult, 'messages'>[]>([]);
    const [selection, setSelection] = useState<number | 'new-parse'>('new-parse');

    useEffect(() => {
        // TODO: fetch parses from backend
        setParses(mockData)
    }, [])

    return (
        <AppShell
            padding="md"
            navbar={
                <Sidenav
                    parses={parses}
                    selection={selection}
                    setSelection={(newSelection: number | 'new-parse') => setSelection(newSelection)}
                />
            }
        >
            {
                selection === 'new-parse' ?
                    <NewParse/> :
                    <ParseResult id={selection} />
            }
        </AppShell>
    )
}

export default App
