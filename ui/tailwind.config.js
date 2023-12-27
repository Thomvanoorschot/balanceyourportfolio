/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {
            spacing: {
                '128': '32rem',
            },
            colors: {
                primary: "#f3d2c1",
                secondary: "#fef6e4",
                tertiary: "#001858",
                quaternary: "#f582ae",
            }
        },
    },
    darkMode: 'class',
    plugins: [],
}

