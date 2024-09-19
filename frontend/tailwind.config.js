/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#0F52BA',
        secondary: '#49D4C6',
        darkGray: '#333333'
      },
      maxWidth: {
        'mobile': '100%',
        'tablet': '768px',
        'desktop': '1024px',
      },
    },
  },
  plugins: [],
}
