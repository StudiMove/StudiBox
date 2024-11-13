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
        lightBlue: '#E7EEF9',
        darkGray: '#333333',
        lightBlack: '#2B323B',
        lightGray: '#B8C0CC',
        lightGreen: '#EDFBFA',
        whiteBlack: '#F3F4F6',
        red: '#FF2453'

      },
      maxWidth: {
        'mobile': '100%',
        'tablet': '768px',
        'desktop': '1024px',
      },
      width: {
        '120': '496px',
      },
      fontFamily: {
        helvetica: ['Helvetica Neue', 'Helvetica', 'Arial', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
