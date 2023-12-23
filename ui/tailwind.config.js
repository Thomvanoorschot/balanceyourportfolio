/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        primary: "var(--primary)",
        primaryhighlighted: "var(--primary-highlighted)",
        secondary: "var(--secondary)",
        secondaryhighlighted: "var(--secondary-highlighted)",
      }
    },
  },
  darkMode: 'class',
  plugins: [],
}

