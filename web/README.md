# KPass Client

## Requirements
* node `^4.5.0`
* yarn `^0.17.0` or npm `^3.0.0`

## Development

```bash
$ npm install    # Install project dependencies
$ npm start      # Compile and launch
```

While developing, you will probably rely mostly on `npm start`; however, there are additional scripts at your disposal:

|`npm run <script>`|Description|
|------------------|-----------|
|`start`|Serves your app at `localhost:3000`. HMR will be enabled in development.|
|`compile`|Compiles the application to disk (`~/dist` by default).|
|`dev`|Same as `npm start`, but enables nodemon for the server as well.|
|`test`|Runs unit tests with Karma and generates a coverage report.|
|`test:dev`|Runs Karma and watches for changes to re-run tests; does not generate coverage reports.|
|`deploy`|Runs linter, tests, and then, on success, compiles your application to disk.|
|`deploy:dev`|Same as `deploy` but overrides `NODE_ENV` to "development".|
|`deploy:prod`|Same as `deploy` but overrides `NODE_ENV` to "production".|
|`lint`|Lint all `.js` files.|
|`lint:fix`|Lint and fix all `.js` files. [Read more on this](http://eslint.org/docs/user-guide/command-line-interface.html#fix).|

## Application Structure

```
.
├── bin                      # Build/Start scripts
├── config                   # Project and build configurations
├── public                   # Static public assets (not imported anywhere in source code)
├── server                   # Express application that provides webpack middleware
│   └── main.js              # Server application entry point
├── src                      # Application source code
│   ├── index.html           # Main HTML page container for app
│   ├── main.js              # Application bootstrap and rendering
│   ├── routes               # Main route definitions and async split points
│   │   ├── index.js         # Bootstrap main application routes with store
│   │   └── example          # Fractal route
│   │       ├── index.js     # Counter route definition
│   │       ├── container.js # Connect components to actions and store
│   │       ├── assets       # Assets required to render components
│   │       ├── modules      # Collections of reducers/constants/actions
│   │       ├── views        # Presentational React View Components
│   │       └── routes **    # Fractal sub-routes (** optional)
│   ├── store                # Redux-specific pieces
│   │   └── reducers         # Reducer registry and injection
│   │   ├── createStore.js   # Create and instrument redux store
│   │   ├── initialState.js  # Redux store's initial state
│   └── styles               # Application-wide styles (generally settings)
│   ├── uis                  # Global Reusable UI Components
│   ├── utils                # Global Utilities
│   ├── views                # Global Reusable Presentational View Components
└── tests                    # Unit tests
```

The application structure is **fractal**, If you wish to read more about this pattern, please check out this [awesome writeup](https://github.com/davezuko/react-redux-starter-kit/wiki/Fractal-Project-Structure) by [Justin Greenberg](https://github.com/justingreenberg).

## Testing
To add a unit test, simply create a `.spec.js` file anywhere in `~/tests`. Karma will pick up on these files automatically, and Mocha and Chai will be available within your test without the need to import them. Coverage reports will be compiled to `~/coverage` by default. If you wish to change what reporters are used and where reports are compiled, you can do so by modifying `coverage_reporters` in `~/config/project.config.js`.

## Build System

### Configuration

Default project configuration can be found in `~/config/project.config.js`. Here you'll be able to redefine your `src` and `dist` directories, adjust compilation settings, tweak your vendor dependencies, and more. For the most part, you should be able to make changes in here **without ever having to touch the actual webpack build configuration**.

If you need environment-specific overrides (useful for dynamically setting API endpoints, for example), you can edit `~/config/environments.config.js` and define overrides on a per-NODE_ENV basis. There are examples for both `development` and `production`, so use those as guidelines. Here are some common configuration options:

|Key|Description|
|---|-----------|
|`dir_src`|application source code base path|
|`dir_dist`|path to build compiled application to|
|`server_host`|hostname for the Express server|
|`server_port`|port for the Express server|
|`compiler_devtool`|what type of source-maps to generate (set to `false`/`null` to disable)|
|`compiler_vendor`|packages to separate into to the vendor bundle|

Webpack is configured to make use of [resolve.root](http://webpack.github.io/docs/configuration.html#resolve-root), which lets you import local packages as if you were traversing from the root of your `~/src` directory. Here's an example:

```js
// current file: ~/src/views/some/nested/View.js
// What used to be this:
import SomeComponent from '../../../components/SomeComponent'

// Can now be this:
import SomeComponent from 'components/SomeComponent' // Hooray!
```

### Globals

These are global variables available to you anywhere in your source code. If you wish to modify them, they can be found as the `globals` key in `~/config/project.config.js`. When adding new globals, make sure you also add them to `~/.eslintrc`.

|Variable|Description|
|---|---|
|`process.env.NODE_ENV`|the active `NODE_ENV` when the build started|
|`__DEV__`|True when `process.env.NODE_ENV` is `development`|
|`__PROD__`|True when `process.env.NODE_ENV` is `production`|
|`__TEST__`|True when `process.env.NODE_ENV` is `test`|

### Styles

Both `.styl` and `.css` file extensions are supported out of the box. After being imported, styles will be processed with [PostCSS](https://github.com/postcss/postcss) for minification and autoprefixing, and will be extracted to a `.css` file during production builds.

### Server

This starter kit comes packaged with an Express server. It's important to note that the sole purpose of this server is to provide `webpack-dev-middleware` and `webpack-hot-middleware` for hot module replacement. Using a custom Express app in place of [webpack-dev-server](https://github.com/webpack/webpack-dev-server) makes it easier to extend the starter kit to include functionality such as API's, universal rendering, and more -- all without bloating the base boilerplate.

### Production Optimization

Babel is configured to use [babel-plugin-transform-runtime](https://www.npmjs.com/package/babel-plugin-transform-runtime) so transforms aren't inlined. In production, webpack will extract styles to a `.css` file, minify your JavaScript, and perform additional optimizations such as module deduplication.

## Dependences

### Seed
- [react-redux-starter-kit](https://github.com/davezuko/react-redux-starter-kit) - Get started with React, Redux, and React-Router.

### Library
- [react](https://github.com/facebook/react) - A declarative, efficient, and flexible JavaScript library for building user interfaces.
- [redux](https://github.com/reactjs/redux) - Predictable state container for JavaScript apps.
