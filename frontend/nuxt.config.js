import colors from 'vuetify/es5/util/colors'

let DEV = process.env.NODE_ENV !== 'production'

export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: 'MINAWatch',
    title: 'MINAWatch',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: 'Minaview - an ecosystem overview of the Mina Protocol',
      },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    {
      src: './plugins/amChart',
      ssr: false,
    },
  ],
  target: 'static',
  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
    '@nuxtjs/axios',
    'vue-masonry-css',
  ],

  publicRuntimeConfig: {
    recaptcha: {
      /* reCAPTCHA options */
      // hideBadge: Boolean, // Hide badge element (v3 & v2 via size=invisible)
      // language: String,   // Recaptcha language (v2)
      version: 3,
      siteKey: '6LeIakEcAAAAALiVrDKecxEtXLzyN7OQRbCQRKuO', //process.env.RECAPTCHA_SITE_KEY // for example
    },
  },

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: ['@nuxtjs/axios', 'vue-masonry-css', '@nuxtjs/recaptcha'],

  // recaptcha: {
  //   hideBadge: Boolean, // Hide badge element (v3 & v2 via size=invisible)
  //   language: String, // Recaptcha language (v2)
  //   siteKey: String, // Site key for requests
  //   version: Number, // Version
  //   size: String, // Size: 'compact', 'normal', 'invisible' (v2)
  // },

  axios: {
    baseURL: DEV ? 'http://localhost:8000' : 'https://api.minaview.com',
  },
  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3,
        },
      },
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
}
