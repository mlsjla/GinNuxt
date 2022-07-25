module.exports = {
    content: [
        "./app.vue",
        "./components/**/*.{vue,js}",
        "./layouts/**/*.vue",
        "./pages/**/*.vue",
        "./plugins/**/*.{js,ts}",
    ],
    theme: {
        colors: {
            transparent: 'transparent',
            current: '#0054a5',
            'primary': '#0054a5',
            'white': '#ffffff',
            'purple': '#3f3cbb',
            'midnight': '#121063',
            'metal': '#565584',
            'tahiti': '#3ab7bf',
            'silver': '#ecebff',
            'bubble-gum': '#ff77e9',
            'bermuda': '#78dcca',
        },
        extend: {}
    },
    plugins: [
        require('daisyui'),
        require('@tailwindcss/typography'),
    ],
    daisyui: {
        themes: [{
                'mytheme': { // custom theme
                    'primary': '#0054a5',
                    'primary-focus': '#0054a5',
                    'primary-content': '#ffffff',
                    "neutral": "#3D4451",
                    "accent": "#37CDBE",
                    "base-100": "#f6f6f6",
                    // other colors
                }
            }
        ],
        styled: true,
        base: true,
        utils: true,
        logs: false,
        rtl: false,
    },
}