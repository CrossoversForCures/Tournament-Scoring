import flowbitePlugin from 'flowbite/plugin'

import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}', './node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}'],
	darkMode: 'class',
	theme: {
		extend: {
			colors: {
				// flowbite-svelte
				primary: {
					50: '#FFF5F2',
					100: '#FFF1EE',
					200: '#FFE4DE',
					300: '#FFD5CC',
					400: '#FFBCAD',
					500: '#FE795D',
					600: '#EF562F',
					700: '#EB4F27',
					800: '#CC4522',
					900: '#A5371B'
				},
				theme: '#8E43F0',
				hover: '#6300E2',
				heading: '#150E1F',
				text: '#584D66',
				gold: '#FFD700',
				silver: '#C0C0C0',
				bronze: '#CD7F32'
			},
			fontFamily: {
				'default': ['Source Sans Pro', 'sans-serif'],
				'heading': ['Poppins', 'sans-serif'],
			},
		}
	},

	plugins: [flowbitePlugin]

} as Config;