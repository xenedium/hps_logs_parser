import React from 'react';
import {ColorScheme, ColorSchemeProvider, MantineProvider} from '@mantine/core';
import App from './App.tsx';

export const Index = () => {
    const [colorScheme, setColorScheme] = React.useState<ColorScheme>('light')
    const toggleColorScheme = (value?: ColorScheme) =>
        setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

    return (
        <React.StrictMode>
            <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
                <MantineProvider theme={{ colorScheme }} withGlobalStyles withNormalizeCSS>
                    <App />
                </MantineProvider>
            </ColorSchemeProvider>
        </React.StrictMode>
    )
}