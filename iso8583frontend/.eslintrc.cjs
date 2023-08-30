module.exports = {
    root: true,
    env: {browser: true, es2020: true},
    extends: [
        'eslint:recommended',
        'plugin:@typescript-eslint/recommended',
        'plugin:react-hooks/recommended',
    ],
    ignorePatterns: ['dist', '.eslintrc.cjs'],
    parser: '@typescript-eslint/parser',
    plugins: ['react-refresh'],
    rules: {
        'react-refresh/only-export-components': [
            'warn',
            {allowConstantExport: true},
        ],
        'indent': ['error', 4, {"ImportDeclaration": 1, "ObjectExpression": 1, "ArrayExpression": 1}],
        'block-spacing': ['error', 'always'],
        'arrow-spacing': ['error', {before: true, after: true}],
        'comma-spacing': ['error', {before: false, after: true}],
        'array-bracket-spacing': ['error', 'never'],
        'func-call-spacing': ['error', 'never'],
        'comma-style': ['error', 'last'],
        'quotes': ['error', 'single'],
    },
}
