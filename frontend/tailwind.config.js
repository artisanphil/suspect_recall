/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}", // Scan your React project files for Tailwind classes
  ],
  theme: {
    extend: {
      fontFamily: {
        'wild-west': ['Special Elite', 'serif'],
      },
    }, 
  },
  plugins: [],
};
