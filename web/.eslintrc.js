// https://eslint.org/docs/user-guide/configuring

module.exports = {
	root: true,
	parserOptions: {
		parser: 'babel-eslint',
		sourceType: 'module',
	},
	parser: 'vue-eslint-parser',
	env: {
		browser: true,
		node: true,
		es6: true,
	},
	// required to lint *.vue files
	plugins: ['vue'],
	// add your custom rules here
	rules: {
		// allow async-await
		'generator-star-spacing': 'off',
		'key-spacing': 'off',
		'no-console': 'off',

		// allow debugger during development
		'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
	},
}
