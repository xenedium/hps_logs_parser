import {AppShell} from '@mantine/core';

import {Sidenav} from './components/Sidenav.tsx';
import {useEffect, useState} from 'react';
import NewParse from './pages/NewParse.tsx';
import ParseResult from './pages/ParseResult.tsx';

function App() {
    const [parses, setParses] = useState<string[]>([]);
    const [selection, setSelection] = useState<string | 'new-parse'>('new-parse');

    useEffect(() => {
        fetch(`${import.meta.env.DEV ? 'http://127.0.0.1:8000' : ''}/api/v1/keys`)
            .then(response => response.json())
            .then(data => setParses(data.keys))
    }, [])

    return (
        <AppShell
            padding="md"
            navbar={
                <Sidenav
                    parses={parses}
                    selection={selection}
                    setSelection={(newSelection: string | 'new-parse') => setSelection(newSelection)}
                />
            }
        >
            {
                selection === 'new-parse' ?
                    <NewParse/> :
                    <ParseResult name={selection} />
            }
        </AppShell>
    )
}

export default App
